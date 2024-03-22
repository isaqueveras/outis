package outis

// Context ...
type Context struct {
	Id           *ID
	Name         *string
	Desc         *string
	StartHour    *int
	EndHour      *int
	Interval     *int
	LoadInterval *int

	routine Routine
	channel chan interface{}

	path   string
	metric *Metric
}

// GetID ...
func (ctx *Context) GetID() ID {
	if ctx.Id != nil {
		return *ctx.Id
	}
	return ID("id")
}

// Error ...
func (ctx *Context) Error(e error) {
	ctx.channel <- e
}

// Info ...
func (ctx *Context) Info(msg string) {
	ctx.channel <- msg
}

// GetName ...
func (ctx *Context) GetName() string {
	if ctx.Name != nil {
		return *ctx.Name
	}
	return "no name"
}

// Metric ...
func (ctx *Context) Metric(value string, args interface{}) {
	ctx.metric.Set(value, args)
}

func (o *Context) isTime(hour int) bool {
	if o.StartHour == nil || o.EndHour == nil {
		return true
	}

	if *o.StartHour == 0 && *o.EndHour == 0 {
		return true
	}

	if *o.StartHour <= *o.EndHour {
		return (hour >= *o.StartHour && hour <= *o.EndHour)
	}

	return (hour >= *o.StartHour || hour <= *o.EndHour)
}
