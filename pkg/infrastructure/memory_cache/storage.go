package memory_cache

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"sync"
)

type BankStorage struct {
	list map[int]balance
	m    sync.RWMutex
}

func (s *BankStorage) CreateAccount(user *services.CreateAccount) *entity.User {
	userID := user.UserID
	b := balance{
		b: user.Balance,
	}

	s.m.Lock()
	defer s.m.Unlock()

	s.list[userID] = b
	return b.ToEntity(userID)
}

func (s *BankStorage) GetBalance(user *services.GetBalance) *services.GetBalanceResponse {
	//TODO implement me
	panic("implement me")
}

func (s *BankStorage) ListAccounts() *services.ListAccountResponse {
	//TODO implement me
	panic("implement me")
}

func (s *BankStorage) UpdateBalance(user *services.UpdateBalance) *services.UpdateBalanceResponse {
	//TODO implement me
	panic("implement me")
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
