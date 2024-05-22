package outis

import (
	"sync"
	"time"
)

var metricMutex sync.Mutex

type Metric map[string]interface{}

func (metric Metric) Set(value string, args interface{}) {
	metricMutex.Lock()
	defer metricMutex.Unlock()
	metric[value] = args
}

type EventMetric struct {
	ID         string        `json:"id"`
	Latency    time.Duration `json:"latency"`
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at"`
	Watcher    WatcherMetric `json:"watcher"`
	Routine    RoutineMetric `json:"routine"`
	Metadata   Metric        `json:"metadata"`
	Indicator  []Indicator   `json:"indicators"`
	Log        []Log         `json:"logs"`
}

type RoutineMetric struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	StartedAt time.Time `json:"started_at"`
}

type WatcherMetric struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartedAt time.Time `json:"started_at"`
}
