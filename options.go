package outis

import "time"

// With ...
func (ctx *Context) With(opts ...Option) {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(ctx)
	}
}

// WithName ...
func WithName(name string) Option {
	return func(r *Context) {
		r.name = name
	}
}

// WithDesc ...
func WithDesc(desc string) Option {
	return func(r *Context) {
		r.desc = desc
	}
}

// WithID ...
func WithID(id ID) Option {
	return func(r *Context) {
		r.id = id
	}
}

// WithRoutine ...
func WithRoutine(routine Routine) Option {
	return func(r *Context) {
		r.routine = routine
	}
}

// WithHours ...
func WithHours(start, end uint) Option {
	return func(r *Context) {
		r.startHour, r.endHour = start, end
	}
}

// WithInterval ...
func WithInterval(duration time.Duration) Option {
	return func(r *Context) {
		r.interval = duration
	}
}

// WithLoadInterval ...
func WithLoadInterval(duration time.Duration) Option {
	return func(r *Context) {
		r.loadInterval = duration
	}
}
