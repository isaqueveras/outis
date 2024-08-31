package outis

import (
	"errors"
	"time"
)

// Context defines the data structure of the routine context
type Context struct {
	Id        ID            `json:"id,omitempty"`
	RoutineID ID            `json:"routine_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Desc      string        `json:"desc,omitempty"`
	Start     uint          `json:"start,omitempty"`
	End       uint          `json:"end,omitempty"`
	Interval  time.Duration `json:"interval,omitempty"`
	Path      string        `json:"path,omitempty"`
	RunAt     time.Time     `json:"run_at,omitempty"`
	Watcher   Watch         `json:"-"`

	script     func(*Context) error
	metadata   Metadata
	latency    time.Duration
	histrogram []*histogram
	indicator  []*indicator
	log        ILogger
}

// GetLatency get script execution latency (in seconds)
func (ctx *Context) GetLatency() float64 {
	return ctx.latency.Seconds()
}

// Error creates a new error message
func (ctx *Context) Error(msg string, v ...interface{}) {
	ctx.log.Errorf(msg, v...)
}

// Info creates a new info message
func (ctx *Context) Info(msg string, v ...interface{}) {
	ctx.log.Infof(msg, v...)
}

// Debug creates a new debug message
func (ctx *Context) Debug(msg string, v ...interface{}) {
	ctx.log.Debugf(msg, v...)
}

// Warn creates a new warn message
func (ctx *Context) Warn(msg string, v ...interface{}) {
	ctx.log.Warnf(msg, v...)
}

// Panic creates a new panic message
func (ctx *Context) Panic(msg string, v ...interface{}) {
	ctx.log.Panicf(msg, v...)
}

// Metadata method for adding data to routine metadata
func (ctx *Context) Metadata(key string, args interface{}) {
	ctx.metadata.Set(key, args)
}

func (ctx *Context) metrics(w *Watch, now time.Time) {
	w.outis.Event(ctx, EventMetric{
		ID:         ctx.Id.ToString(),
		StartedAt:  now,
		FinishedAt: time.Now(),
		Latency:    time.Since(now),
		Metadata:   ctx.metadata,
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

	ctx.metadata, ctx.indicator, ctx.histrogram = Metadata{}, []*indicator{}, []*histogram{}
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
