package outis

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

// Channel ...
type Channel chan error

// Routine ...
type Routine func(Channel)

// Option ...
type Option func(*Outis)

// IOutis ...
type IOutis interface {
	// Lock ...
	Lock(string) (string, error)

	// Unlock ...
	Unlock(string) error

	// Store ...
	Store(*Outis) error

	// Load ...
	Load(*Outis) error

	// Handling ...
	Handling(error)
}

// Outis ...
type Outis struct {
	// Name
	Name string `json:"name"`

	// Description
	Description string `json:"desc"`

	// TypeRoutine
	TypeRoutine string `json:"type"`

	// Path
	Path string `json:"path"`

	// StartHour
	StartHour int `json:"start_date"`

	// EndHour
	EndHour int `json:"end_date"`

	// Interval
	Interval int `json:"interval"`

	routine Routine `json:"-"`
	channel Channel `json:"-"`
}

// Watch ...
type Watch struct {
	ioutis IOutis

	channel Channel
	signal  chan os.Signal
}

// Watcher ...
func Watcher(ioutis IOutis) *Watch {
	w := &Watch{
		channel: make(chan error, 15),
		signal:  make(chan os.Signal, 1),
		ioutis:  ioutis,
	}

	signal.Notify(w.signal, syscall.SIGINT, syscall.SIGTERM)
	return w
}

// Wait ...
func (w *Watch) Wait() {
	for {
		select {
		case <-w.signal:
			log.Println("outis: sinal de encerramento recebido")
			return
		case err := <-w.channel:
			w.ioutis.Handling(err)
		}
	}
}

// Go ...
func (w *Watch) Go(opts ...Option) {
	o := &Outis{Interval: 600} // 60 sec * 10 min
	defer o.recover()

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(o)
	}

	info := runtime.FuncForPC(reflect.ValueOf(o.routine).Pointer())
	file, line := info.FileLine(info.Entry())
	o.Path = fmt.Sprintf("%s:%v", file, line)

	if err := w.ioutis.Store(o); err != nil {
		w.channel <- err
		return
	}

	for range time.NewTicker(time.Second * time.Duration(o.Interval)).C {
		now := time.Now()
		if !o.ItTime(now.Hour()) {
			continue
		}

		w.process(o)
	}
}

func (w *Watch) process(outis *Outis) error {
	if err := w.ioutis.Load(outis); err != nil {
		return err
	}

	id, err := w.ioutis.Lock(outis.TypeRoutine)
	if err != nil {
		return err
	}

	outis.routine(w.channel)

	if err = w.ioutis.Unlock(id); err != nil {
		return err
	}

	return nil
}

func (r *Outis) recover() {
	if rec := recover(); rec != nil {
		const msg = "[script] - [PANIC] "
		if v, ok := rec.(error); ok {
			r.channel <- fmt.Errorf(msg+": %v", v.Error())
		} else {
			r.channel <- fmt.Errorf(msg+": %v", rec)
		}
	}
}

// ItTime ...
func (o *Outis) ItTime(hour int) bool {
	if o.StartHour == 0 && o.EndHour == 0 {
		return true
	}

	if o.StartHour <= o.EndHour {
		return (hour >= o.StartHour && hour <= o.EndHour)
	}

	return (hour >= o.StartHour || hour <= o.EndHour)
}
