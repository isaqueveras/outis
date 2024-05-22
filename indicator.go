package outis

import "time"

type indicator struct {
	Key       string    `json:"key"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func (ctx *Context) NewIndicator(key string) *indicator {
	indicator := &indicator{Key: key, Value: 0, CreatedAt: time.Now()}
	ctx.indicator = append(ctx.indicator, indicator)
	return indicator
}

func (i *indicator) Inc()              { i.Value++ }
func (i *indicator) Add(value float64) { i.Value += value }
