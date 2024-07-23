package memory_cache

import "sync"

type BankStorage struct {
	list map[int]float64
	m    sync.RWMutex
}

func New() *BankStorage {
	return &BankStorage{
		list: make(map[int]float64),
	}
}

func (s *BankStorage) Create(userID int, balance float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.list[userID] = balance
}

func (s *BankStorage) Get(userID int) float64 {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.list[userID]
}
