package outis

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Script func(*Context)

type Context struct {
	Id           ID            `json:"id,omitempty"`
	RoutineID    ID            `json:"routine_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Desc         string        `json:"desc,omitempty"`
	Start        uint          `json:"start,omitempty"`
	End          uint          `json:"end,omitempty"`
	Interval     time.Duration `json:"interval,omitempty"`
	LoadInterval time.Duration `json:"load_interval,omitempty"`
	Path         string        `json:"path,omitempty"`
	RunAt        time.Time     `json:"run_at,omitempty"`
	Watcher      Watch         `json:"-"`

	script   Script
	metadata Metadata
	logs     []Log

	histrogram []*histogram
	indicator  []*indicator

	// L define the log layer interface
	L Logger `json:"-"`

	context.Context
}

func (ctx *Context) Error(message string, args ...interface{}) {
	ctx.L.Errorf(message, args...)
	ctx.logs = append(ctx.logs, Log{
		Message:   fmt.Sprintf(message, args...),
		Level:     levelLogError,
		Timestamp: time.Now(),
	})
}

func (ctx *Context) Info(message string, args ...interface{}) {
	ctx.L.Infof(message, args...)
	ctx.logs = append(ctx.logs, Log{
		Message:   fmt.Sprintf(message, args...),
		Level:     levelLogInfo,
		Timestamp: time.Now(),
	})
}

func (ctx *Context) Debug(message string, args ...interface{}) {
	ctx.L.Debugf(message, args...)
	ctx.logs = append(ctx.logs, Log{
		Message:   fmt.Sprintf(message, args...),
		Level:     levelLogDebug,
		Timestamp: time.Now(),
	})
}

func (ctx *Context) Panic(message string, args ...interface{}) {
	ctx.L.Panicf(message, args...)
	ctx.logs = append(ctx.logs, Log{
		Message:   fmt.Sprintf(message, args...),
		Level:     levelLogPanic,
		Timestamp: time.Now(),
	})
}

func (ctx *Context) Metadata(key string, args interface{}) {
	ctx.metadata.Set(key, args)
}

func (ctx *Context) reload(ioutis Outis) {
	ticker := time.NewTicker(ctx.LoadInterval)
	for range ticker.C {
		if err := ioutis.Reload(ctx); err != nil {
			ctx.L.Errorf(err.Error())
		}

		ticker.Reset(ctx.LoadInterval)
		ctx.Info("script '%s' (rid: %s) has been updated", ctx.Name, ctx.RoutineID)
	}
}

func (ctx *Context) metrics(w *Watch, now time.Time) {
	w.outis.Event(ctx, EventMetric{
		ID:         ctx.Id.ToString(),
		StartedAt:  now,
		FinishedAt: time.Now(),
		Latency:    time.Since(now),
		Metadata:   ctx.metadata,
		Log:        ctx.logs,
		Indicator:  ctx.indicator,
		Histogram:  ctx.histrogram,
		Watcher: WatcherMetric{
			ID:    w.Id.ToString(),
			Name:  w.Name,
			RunAt: w.RunAt,
		},
		Routine: RoutineMetric{
			ID:        ctx.RoutineID.ToString(),
			Name:      ctx.Name,
			Path:      ctx.Path,
			StartedAt: ctx.RunAt,
		},
	})

	ctx.logs, ctx.metadata, ctx.indicator, ctx.histrogram =
		[]Log{}, Metadata{}, []*indicator{}, []*histogram{}
}

func (ctx *Context) isTime(hour int) bool {
	if ctx.Start == 0 && ctx.End == 0 {
		return true
	}

	if ctx.Start <= ctx.End {
		return (hour >= int(ctx.Start) && hour <= int(ctx.End))
	}

	return (hour >= int(ctx.Start) || hour <= int(ctx.End))
}

func (ctx *Context) validate() error {
	if ctx.RoutineID == "" {
		return errors.New("the routine id is required")
	}

	if ctx.Name == "" {
		return errors.New("the routine name is required")
	}

	if ctx.script == nil {
		return errors.New("the routine is required")
	}

	return nil
}
