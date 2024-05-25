package outis

import "time"

type Option func(*Context)

func WithName(name string) Option {
	return func(r *Context) {
		r.Name = name
	}
}

func WithDesc(desc string) Option {
	return func(r *Context) {
		r.Desc = desc
	}
}

func WithID(id ID) Option {
	return func(r *Context) {
		r.RoutineID = id
	}
}

func WithScript(routine Script) Option {
	return func(r *Context) {
		r.script = routine
	}
}

func WithHours(start, end uint) Option {
	return func(r *Context) {
		r.Start, r.End = start, end
	}
}

func WithInterval(duration time.Duration) Option {
	return func(r *Context) {
		r.Interval = duration
	}
}

func WithLoadInterval(duration time.Duration) Option {
	return func(r *Context) {
		r.LoadInterval = duration
	}
}

type WatcherOption func(*Watch)

func WithLogger(logger Logger) WatcherOption {
	return func(r *Watch) { r.log = logger }
}

func WithOutisInterface(outis Outis) WatcherOption {
	return func(r *Watch) { r.outis = outis }
}
