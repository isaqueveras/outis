package outis

type Event interface{}

// Outis is the main interface for implementing the outis lib.
type Outis interface {
	Init(ctx *Context) error
	Before(ctx *Context) error
	After(ctx *Context) error
	Reload(ctx *Context) error
	Event(ctx *Context, event Event)
}

// Logger methods for logging messages.
type Logger interface {
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}
