package outis

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

type ID string

func (id ID) ToString() string {
	return string(id)
}

type Watch struct {
	id        ID
	name      string
	signal    chan os.Signal
	startedAt time.Time

	outis Outis
	log   Logger
}

func Watcher(id, name string, opts ...WatcherOption) *Watch {
	watch := &Watch{
		id:        ID(id),
		name:      name,
		signal:    make(chan os.Signal, 1),
		log:       setupLogger(),
		outis:     newOutis(),
		startedAt: time.Now(),
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(watch)
	}

	signal.Notify(watch.signal, syscall.SIGINT, syscall.SIGTERM)
	return watch
}

func (w *Watch) Wait() {
	for range w.signal {
		w.log.Infof("closing signal received")
		break
	}
}

func (w *Watch) Go(opts ...Option) {
	ctx := &Context{
		indicator:    make([]Indicator, 0),
		metric:       make(Metric),
		logs:         make([]Log, 0),
		LoadInterval: 0,
		L:            w.log,
		Interval:     time.Minute,
		StartedAt:    time.Now(),
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(ctx)
	}

	if err := ctx.validate(); err != nil {
		ctx.Error("%v", err)
		return
	}

	if err := w.outis.Init(ctx); err != nil {
		ctx.Error("%v", err)
		return
	}

	info := runtime.FuncForPC(reflect.ValueOf(ctx.script).Pointer())
	file, line := info.FileLine(info.Entry())
	ctx.Path = fmt.Sprintf("%s:%v", file, line)

	if ctx.LoadInterval != 0 {
		go ctx.reload(w.outis)
	}

	defer func() {
		if r := recover(); r != nil {
			ctx.Error("%v", r)
		}
	}()

	ticker := time.NewTicker(ctx.Interval)
	for range ticker.C {
		if !ctx.isTime(time.Now().Hour()) {
			continue
		}

		now := time.Now()
		if err := w.outis.Before(ctx); err != nil {
			ctx.Error(err.Error())
			continue
		}

		func(script Script) {
			defer func() {
				if err := recover(); err != nil {
					ctx.Error("%v", err)
				}
			}()
			script(ctx)
		}(ctx.script)

		if err := w.outis.After(ctx); err != nil {
			ctx.Error(err.Error())
			continue
		}

		ctx.metrics(w, now)
		ticker.Reset(ctx.Interval)
	}
}
