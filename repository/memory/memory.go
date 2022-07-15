package memory

import (
	"errors"
	"sync"
)

type memory struct {
	kv map[string]interface{}

	mu sync.RWMutex
}

func Init() (m memory) {
	m = memory{
		kv: make(map[string]interface{}),
		mu: sync.RWMutex{},
	}
	return
}

func (m *memory) Get(key string) (val interface{}, err error) {
	m.mu.RLock()
	val, ok := m.kv[key]
	m.mu.RUnlock()

	if !ok {
		err = errors.New("invalid key")
	}
	return
}

func (m *memory) Set(key string, val interface{}) (err error) {
	if m.kv == nil {
		m.mu.Lock()
		m.kv = make(map[string]interface{})
		m.mu.Unlock()
	}

	m.mu.Lock()
	m.kv[key] = val
	m.mu.Unlock()

	return
}
