package outis

import (
	"encoding/json"
	"time"
)

// Metadata defines type of metadata used in metrics
type Metadata map[string]interface{}

// Set add data to metadata
func (md Metadata) Set(value string, args interface{}) {
	md[value] = args
}

// GetBytes return the routine metadata in bytes
func (md Metadata) GetBytes() []byte {
	value, _ := json.Marshal(md)
	return value
}

// EventMetric defines the type of metric sent in the event
type EventMetric struct {
	ID         string
	Latency    time.Duration
	StartedAt  time.Time
	FinishedAt time.Time
	Watcher    WatcherMetric
	Routine    RoutineMetric
	Metadata   Metadata
	Indicators []*indicator
	Histograms []*histogram
}

// RoutineMetric defines the type of metric
// of a routine sent in the event
type RoutineMetric struct {
	ID        string
	Name      string
	Path      string
	StartedAt time.Time
}

// WatcherMetric defines the type of metric
// of a watcher sent in the event
type WatcherMetric struct {
	ID    string
	Name  string
	RunAt time.Time
}
