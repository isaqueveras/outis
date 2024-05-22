package outis

import (
	"errors"
	"fmt"
	"time"
)

type Script func(*Context)

type Context struct {
	Id           ID
	RoutineID    ID
	Name         string
	Desc         string
	StartHour    uint
	EndHour      uint
	Interval     time.Duration
	LoadInterval time.Duration
	Path         string
	StartedAt    time.Time

	script    Script
	metric    Metric
	indicator []Indicator
	logs      []Log

	// L define the log layer interface
	L Logger
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

func (ctx *Context) Metric(key string, args interface{}) {
	ctx.metric.Set(key, args)
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
	latency := time.Since(now)
	if latency > time.Minute {
		latency = latency.Truncate(time.Second)
	}

	w.outis.Event(ctx, EventMetric{
		ID:         ctx.Id.ToString(),
		StartedAt:  now,
		FinishedAt: time.Now(),
		Latency:    latency,
		Metadata:   ctx.metric,
		Log:        ctx.logs,
		Indicator:  ctx.indicator,
		Watcher: WatcherMetric{
			ID:        w.id.ToString(),
			Name:      w.name,
			StartedAt: w.startedAt,
		},
		Routine: RoutineMetric{
			ID:        ctx.RoutineID.ToString(),
			Name:      ctx.Name,
			Path:      ctx.Path,
			StartedAt: ctx.StartedAt,
		},
	})

	ctx.logs, ctx.metric, ctx.indicator = []Log{}, Metric{}, []Indicator{}
}

func (ctx *Context) isTime(hour int) bool {
	if ctx.StartHour == 0 && ctx.EndHour == 0 {
		return true
	}

	if ctx.StartHour <= ctx.EndHour {
		return (hour >= int(ctx.StartHour) && hour <= int(ctx.EndHour))
	}

	return (hour >= int(ctx.StartHour) || hour <= int(ctx.EndHour))
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
