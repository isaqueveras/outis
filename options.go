package outis

import "time"

// Option defines the option type of a routine
type Option func(*Context)

// WithName defines the name of a routine
func WithName(name string) Option {
	return func(ctx *Context) { ctx.Name = name }
}

// WithDesc defines the description of a routine
func WithDesc(desc string) Option {
	return func(ctx *Context) { ctx.Desc = desc }
}

// WithID defines a routine's identifier
func WithID(id ID) Option {
	return func(ctx *Context) { ctx.RoutineID = id }
}

// WithScript defines the script function that will be executed
func WithScript(fn func(*Context) error) Option {
	return func(ctx *Context) { ctx.script = fn }
}

// WithHours sets the start and end time of script execution
func WithHours(start, end uint) Option {
	return func(ctx *Context) { ctx.period.startHour, ctx.period.endHour = start, end }
}

// WithMinutes sets the start and end minutes of script execution
func WithMinutes(start, end uint) Option {
	return func(ctx *Context) { ctx.period.startMinute, ctx.period.endMinute = start, end }
}

// WithInterval defines the interval at which the script will be executed
func WithInterval(duration time.Duration) Option {
	return func(ctx *Context) { ctx.Interval = duration }
}

// WithNotUseLoop define that the routine will not enter a loop
func WithNotUseLoop() Option {
	return func(ctx *Context) { ctx.notUseLoop = true }
}

// WatcherOption defines the option type of a watcher
type WatcherOption func(*Watch)

// Logger defines the implementation of the log interface
func Logger(logger ILogger) WatcherOption {
	return func(watch *Watch) { watch.log = logger }
}

// Impl defines the implementation of the main interface
func Impl(outis IOutis) WatcherOption {
	return func(watch *Watch) { watch.outis = outis }
}
