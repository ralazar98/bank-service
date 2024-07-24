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

func (s *BankStorage) Create(userID int, balance float64) error {
	s.m.Lock()
	defer s.m.Unlock()
	s.list[userID] = balance
	return nil
}

func (s *BankStorage) Get(userID int) (float64, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.list[userID], nil
}

func (s *BankStorage) List() (map[int]float64, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.list, nil
}

func (s *BankStorage) Update(userID int, changingInBalance float64) error {
	s.m.Lock()
	defer s.m.Unlock()
	s.list[userID] += changingInBalance
	return nil
}
