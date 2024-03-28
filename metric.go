package outis

import (
	"sync"
	"time"
)

// metricMutex ...
var metricMutex sync.Mutex

// Metric
type Metric map[string]interface{}

// Set ...
func (metric Metric) Set(value string, args interface{}) {
	metricMutex.Lock()
	defer metricMutex.Unlock()
	metric[value] = args
}

// Metrics ...
type Metrics struct {
	ID          string        `json:"id"`
	Initialized time.Time     `json:"initialized"`
	Terminated  time.Time     `json:"terminated"`
	Runtime     float64       `json:"runtime"`
	Routine     RoutineMetric `json:"routine"`
	Metadata    Metric        `json:"metadata"`
}

// RoutineMetric ...
type RoutineMetric struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
