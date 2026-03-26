package store

import "time"

func (s *Store) SetWithTTL(key, value string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	s.expiry[key] = time.Now().Add(ttl)
}

func (s *Store) SetExpiry(key string, ttl time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return false // key doesn't exist
	}
	s.expiry[key] = time.Now().Add(ttl)
	return true
}

func (s *Store) TTL(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	exp, ok := s.expiry[key]
	if !ok {
		return -1 // no expiry set
	}
	remaining := int(time.Until(exp).Seconds())
	if remaining <= 0 {
		return -2 // expired
	}
	return remaining
}

func (s *Store) isExpired(key string) bool {
	exp, ok := s.expiry[key]
	if !ok {
		return false
	}
	return time.Now().After(exp)
}

// cleanup runs in background, deletes expired keys every second
func (s *Store) cleanup() {
	for {
		time.Sleep(time.Second)
		s.mu.Lock()
		for key := range s.expiry {
			if time.Now().After(s.expiry[key]) {
				delete(s.data, key)
				delete(s.expiry, key)
			}
		}
		s.mu.Unlock()
	}
}
