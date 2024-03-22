package memory

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/isaqueveras/outis"
)

type memory struct {
	data map[outis.ID]bool
	mu   sync.Mutex
}

// NewOutis ...
func NewOutis() outis.Interface {
	return &memory{data: make(map[outis.ID]bool)}
}

// Lock defines the method by which concurrency
// blocking will be implemented
func (m *memory) Lock(id outis.ID) (outis.ID, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[id] = true
	return outis.ID(uuid.NewString()), nil
}

// Unlock defines the method by which concurrency
// unblocking will be implemented
func (m *memory) Unlock(id outis.ID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[id] = false
	return nil
}

// Store defines the method for saving
// the routine's initial information
func (m *memory) Store(o *outis.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[o.GetID()] = true
	return nil
}

// Load defines the method to fetch updated
// information from the routine
func (m *memory) Load(o *outis.Context) error {
	return nil
}

// Event defines the method for handling events
func (m *memory) Event(event outis.Event) {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch value := event.(type) {
	case error:
		log.Println("outis: [error] " + value.Error())
	case string:
		log.Println("outis: [info] " + value)
	case outis.Metrics:
		v, _ := json.Marshal(value)
		log.Println("outis: [metrics]", string(v))
	}
}
