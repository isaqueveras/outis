package outis

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

type server struct{}

func newOutis() Outis { return &server{} }

func (server) Init(ctx *Context) error {
	ctx.Info("script '%s' (rid: %s) initialized", ctx.Name, ctx.RoutineID)
	return nil
}

func (server) Before(ctx *Context) error {
	ctx.Id = ID(strconv.FormatInt(rand.Int63(), 10))
	ctx.Info("script '%s' (rid: %s, id: %s) initialized", ctx.Name, ctx.RoutineID, ctx.Id)
	return nil
}

func (server) After(ctx *Context) error {
	ctx.Info("script '%s' (rid: %s, id: %s) finished", ctx.Name, ctx.RoutineID, ctx.Id)
	return nil
}

func (server) Reload(ctx *Context) error {
	ctx.Info("script '%s' (rid: %s, id: %s) reloaded", ctx.Name, ctx.RoutineID, ctx.Id)
	return nil
}

func (server) Event(ctx *Context, event Event) {
	if metric, ok := event.(EventMetric); ok {
		value, _ := json.Marshal(metric)
		ctx.Debug(string(value))
	}
}
