package outis

import "time"

type histogram struct {
	key    string
	values []histogramValue
}

type histogramValue struct {
	value     float64
	createdAt time.Time
}

// NewHistogram creates a new histrogram
func (ctx *Context) NewHistogram(key string) *histogram {
	histogram := &histogram{key: key, values: make([]histogramValue, 0)}
	ctx.histrogram = append(ctx.histrogram, histogram)
	return histogram
}

// GetKey get the key value of an histogram
func (h *histogram) GetKey() string { return h.key }

// GetValue get the values of an histogram
func (h *histogram) GetValues() (values []float64, times []time.Time) {
	for _, item := range h.values {
		values, times = append(values, item.value), append(times, item.createdAt)
	}
	return
}

// Inc increments the histogram data
func (h *histogram) Inc() {
	var value float64 = 1
	if len(h.values) != 0 {
		value = h.values[len(h.values)-1].value + 1
	}
	h.values = append(h.values, histogramValue{value: value, createdAt: time.Now()})
}

// Add add a value to the histogram
func (h *histogram) Add(value float64) {
	h.values = append(h.values, histogramValue{value: value, createdAt: time.Now()})
}
