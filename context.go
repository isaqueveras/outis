package outis

import (
	"context"
	"errors"
	"time"
)

type period struct {
	startHour, endHour     uint
	startMinute, endMinute uint
}

// Context defines the data structure of the routine context
type Context struct {
	Id        ID
	RoutineID ID
	Name      string
	Desc      string
	period    period
	Interval  time.Duration
	Path      string
	RunAt     time.Time
	Watcher   Watch

	script     func(*Context) error
	metadata   Metadata
	latency    time.Duration
	notUseLoop bool
	histrogram []*histogram
	indicator  []*indicator
	log        ILogger
	context    context.Context
}

// Deadline returns the time at which work performed on behalf of this context should be canceled.
func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.context.Deadline()
}

// Done returns a channel that's closed when work done on behalf of this context should be canceled.
func (ctx *Context) Done() <-chan struct{} {
	return ctx.context.Done()
}

// Err returns the context error
func (ctx *Context) Err() error {
	return ctx.context.Err()
}

// Value returns the value associated with this context for key
func (ctx *Context) Value(key any) any {
	return ctx.context.Value(key)
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
		Indicators: ctx.indicator,
		Histograms: ctx.histrogram,
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

func (ctx *Context) sleep(now time.Time) {
	if ctx.mustWait(now.Hour(), ctx.period.startHour, ctx.period.endHour) {
		time.Sleep(time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+int(ctx.period.startHour),
			0, 0, 0, now.Location()).Sub(now))
	}

	if ctx.mustWait(now.Minute(), ctx.period.startMinute, ctx.period.endMinute) {
		time.Sleep(time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+int(ctx.period.startHour),
			now.Minute()+int(ctx.period.startMinute), 0, 0, now.Location()).Sub(now))
	}
}

func (ctx *Context) mustWait(time int, start, end uint) bool {
	if start == 0 && end == 0 {
		return false
	}

	if start <= end {
		return !(time >= int(start) && time <= int(end))
	}

	return !(time >= int(start) || time <= int(end))
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
