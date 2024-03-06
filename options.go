package outis

// WithName ...
func WithName(name string) Option {
	return func(r *Outis) {
		r.Name = name
	}
}

// WithDescription ...
func WithDescription(desc string) Option {
	return func(r *Outis) {
		r.Description = desc
	}
}

// WithType ...
func WithType(t string) Option {
	return func(r *Outis) {
		r.TypeRoutine = t
	}
}

// WithRoutine ...
func WithRoutine(routine Routine) Option {
	return func(r *Outis) {
		r.routine = routine
	}
}

// WithHours ...
func WithHours(start, end int) Option {
	return func(r *Outis) {
		r.StartHour, r.EndHour = start, end
	}
}

// WithInterval ...
func WithInterval(seconds int) Option {
	return func(r *Outis) {
		r.Interval = seconds
	}
}
