package outis

// WithName ...
func WithName(name string) Option {
	return func(r *Context) {
		r.Name = &name
	}
}

// WithDesc ...
func WithDesc(desc string) Option {
	return func(r *Context) {
		r.Desc = &desc
	}
}

// WithID ...
func WithID(id ID) Option {
	return func(r *Context) {
		r.Id = &id
	}
}

// WithRoutine ...
func WithRoutine(routine Routine) Option {
	return func(r *Context) {
		r.routine = routine
	}
}

// WithHours ...
func WithHours(start, end int) Option {
	return func(r *Context) {
		r.StartHour, r.EndHour = &start, &end
	}
}

// WithLoadInterval ...
func WithLoadInterval(seconds int) Option {
	return func(r *Context) {
		r.LoadInterval = &seconds
	}
}

// WithInterval ...
func WithInterval(seconds int) Option {
	return func(r *Context) {
		r.Interval = &seconds
	}
}
