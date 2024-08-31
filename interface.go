package outis

// Event sets the type for the event structure
type Event interface{}

// IOutis is the main interface for implementing the outis lib.
type IOutis interface {
	Go(fn func() error)
	Wait() error

	Init(ctx *Context) error
	Before(ctx *Context) error
	After(ctx *Context) error
	Event(ctx *Context, event Event)
}

// ILogger methods for logging messages.
type ILogger interface {
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}
