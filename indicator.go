package outis

import "time"

type indicator struct {
	Key       string    `json:"key"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

// NewIndicator creates a new indicator
func (ctx *Context) NewIndicator(key string) *indicator {
	indicator := &indicator{Key: key, Value: 0, CreatedAt: time.Now()}
	ctx.indicator = append(ctx.indicator, indicator)
	return indicator
}

// Inc increments the indicator data
func (i *indicator) Inc() { i.Value++ }

// Add add a value to the indicator
func (i *indicator) Add(value float64) { i.Value += value }
