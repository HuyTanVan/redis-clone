package store

import (
	"sync"
	"time"
)

type Store struct {
	mu     sync.RWMutex
	data   map[string]string            // SET/GET
	hdata  map[string]map[string]string // HSET/HGET
	expiry map[string]time.Time
}

// Store is an in-memory key-value store that supports basic string operations and hash operations.
func NewStore() *Store {
	s := &Store{
		data:   make(map[string]string),            //data lives in RAM
		hdata:  make(map[string]map[string]string), //hdata lives in RAM
		expiry: make(map[string]time.Time),
	}
	// Start the cleanup goroutine to remove expired keys
	go s.cleanup()
	return s
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
