package outis

import "time"

type histogram struct {
	Key    string           `json:"key"`
	Values []histogramValue `json:"values"`
}

type histogramValue struct {
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func (ctx *Context) NewHistogram(key string) *histogram {
	histogram := &histogram{Key: key, Values: make([]histogramValue, 0)}
	ctx.histrogram = append(ctx.histrogram, histogram)
	return histogram
}

func (i *histogram) Inc() {
	var value float64 = 1
	if len(i.Values) != 0 {
		value = i.Values[len(i.Values)-1].Value + 1
	}
	i.Values = append(i.Values, histogramValue{Value: value, CreatedAt: time.Now()})
}

func (i *histogram) Add(value float64) {
	i.Values = append(i.Values, histogramValue{Value: value, CreatedAt: time.Now()})
}
