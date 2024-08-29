package outis

import "time"

// Option defines the option type of a routine
type Option func(*Context)

// WithName defines the name of a routine
func WithName(name string) Option {
	return func(r *Context) { r.Name = name }
}

// WithDesc defines the description of a routine
func WithDesc(desc string) Option {
	return func(r *Context) { r.Desc = desc }
}

// WithID defines a routine's identifier
func WithID(id ID) Option {
	return func(r *Context) { r.RoutineID = id }
}

// WithScript defines the script function that will be executed
func WithScript(fn func(*Context) error) Option {
	return func(r *Context) { r.script = fn }
}

// WithHours sets the start and end time of script execution
func WithHours(start, end uint) Option {
	return func(r *Context) { r.Start, r.End = start, end }
}

// WithInterval defines the interval at which the script will be executed
func WithInterval(duration time.Duration) Option {
	return func(r *Context) { r.Interval = duration }
}

// WatcherOption defines the option type of a watcher
type WatcherOption func(*Watch)

// Logger defines the implementation of the log interface
func Logger(logger ILogger) WatcherOption {
	return func(r *Watch) { r.log = logger }
}

// Impl defines the implementation of the main interface
func Impl(outis IOutis) WatcherOption {
	return func(r *Watch) { r.outis = outis }
}
