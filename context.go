package outis

// Context ...
type Context struct {
	id           ID
	name         string
	desc         string
	startHour    uint
	endHour      uint
	interval     uint
	loadInterval uint
	path         string
	routine      Routine
	channel      chan interface{}
	metric       Metric
}

// GetID ...
func (ctx *Context) GetID() ID {
	return ctx.id
}

// GetName ...
func (ctx *Context) GetName() string {
	return ctx.name
}

// GetDesc ...
func (ctx *Context) GetDesc() string {
	return ctx.desc
}

// GetStartHour ...
func (ctx *Context) GetStartHour() uint {
	return ctx.startHour
}

// GetEndHour ...
func (ctx *Context) GetEndHour() uint {
	return ctx.endHour
}

// GetPath ...
func (ctx *Context) GetPath() string {
	return ctx.path
}

// GetInterval ...
func (ctx *Context) GetInterval() uint {
	return ctx.interval
}

// GetLoadInterval ...
func (ctx *Context) GetLoadInterval() uint {
	return ctx.loadInterval
}

// Error ...
func (ctx *Context) Error(e error) {
	ctx.channel <- e
}

// Info ...
func (ctx *Context) Info(msg string) {
	ctx.channel <- msg
}

// Metric ...
func (ctx *Context) Metric(value string, args interface{}) {
	ctx.metric.Set(value, args)
}

func (o *Context) isTime(hour int) bool {
	if o.startHour == 0 && o.endHour == 0 {
		return true
	}

	if o.startHour <= o.endHour {
		return (hour >= int(o.startHour) && hour <= int(o.endHour))
	}

	return (hour >= int(o.startHour) || hour <= int(o.endHour))
}
