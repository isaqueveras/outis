package outis

import "time"

type indicator struct {
	key       string
	value     float64
	createdAt time.Time
}

// NewIndicator creates a new indicator
func (ctx *Context) NewIndicator(key string) *indicator {
	indicator := &indicator{key: key, value: 0, createdAt: time.Now()}
	ctx.indicator = append(ctx.indicator, indicator)
	return indicator
}

// GetKey get the key value of an indicator
func (i *indicator) GetKey() string { return i.key }

// GetValue get the value of an indicator
func (i *indicator) GetValue() float64 { return i.value }

// GetCreatedAt get the creation date of an indicator
func (i *indicator) GetCreatedAt() time.Time { return i.createdAt }

// Inc increments the indicator data
func (i *indicator) Inc() { i.value++ }

// Add add a value to the indicator
func (i *indicator) Add(value float64) { i.value += value }
