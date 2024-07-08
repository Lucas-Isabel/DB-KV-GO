package storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) SETk(key, value string) bool {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
	_, includes := s.GETk(key)
	return includes
}

func (s *Storage) GETk(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, found := s.data[key]
	return value, found
}

func (s *Storage) Delete(key string) bool {
	if _, includes := s.GETk(key); !includes {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return true
}

type KV struct {
	KEY   string `json:"key"`
	VALUE string `json:"value"`
}

func (s *Storage) ALLkv() []KV {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var kvs []KV
	for k, v := range s.data {
		kv := KV{
			KEY:   k,
			VALUE: v,
		}
		kvs = append(kvs, kv)
	}
	return kvs
}
