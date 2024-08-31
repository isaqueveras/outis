package outis

import "time"

type Metadata map[string]interface{}

func (metric Metadata) Set(value string, args interface{}) {
	metric[value] = args
}

type EventMetric struct {
	ID         string        `json:"id"`
	Latency    time.Duration `json:"latency"`
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at"`
	Watcher    WatcherMetric `json:"watcher"`
	Routine    RoutineMetric `json:"routine"`
	Metadata   Metadata      `json:"metadata"`
	Indicator  []*indicator  `json:"indicators"`
	Histogram  []*histogram  `json:"histograms"`
}

type RoutineMetric struct {
	ID        string    `json:"routine_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	StartedAt time.Time `json:"started_at"`
}

type WatcherMetric struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	RunAt time.Time `json:"run_at"`
}
