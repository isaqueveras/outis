package outis

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
func WithInterval(seconds uint) Option {
	return func(r *Context) {
		r.interval = seconds
	}
}

// WithLoadInterval ...
func WithLoadInterval(seconds uint) Option {
	return func(r *Context) {
		r.loadInterval = seconds
	}
}
