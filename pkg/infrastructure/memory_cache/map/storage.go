package _map

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"sync"
)

type BankStorage struct {
	list map[int]Balance
	m    sync.RWMutex
}

func New() *BankStorage {
	return &BankStorage{
		list: make(map[int]Balance),
	}
}

func (s *BankStorage) CreateAccount(user *services.CreateAccount) (*entity.User, error) {

	var userRep User
	userRep.ID = user.UserID
	userRep.Balance = Balance{
		Sum: user.Balance,
	}
	_, ok := s.list[user.UserID]
	if ok {
		return nil, services.AccountAlreadyExistsErr
	} else {
		s.m.Lock()
		defer s.m.Unlock()
		s.list[userRep.ID] = userRep.Balance
	}
	return userRep.ToEntity(), nil
}

func (s *BankStorage) GetBalance(user *services.GetBalance) (*entity.User, error) {
	var userRep User
	userRep.ID = user.UserID
	_, ok := s.list[user.UserID]
	if ok {
		s.m.RLock()
		defer s.m.RUnlock()
		userRep.Balance = s.list[userRep.ID]
		return userRep.ToEntity(), nil
	} else {
		return nil, services.ChosenAccountNotFoundErr
	}
}

func (s *BankStorage) UpdateBalance(user *services.UpdateBalance) (*entity.User, error) {
	var userRep User
	s.m.Lock()
	defer s.m.Unlock()
	userRep.ID = user.UserID
	balance, ok := s.list[userRep.ID]
	if ok {
		if balance.Sum+user.ChangingInBalance < 0 {
			return nil, services.NotEnoughBalanceErr
		}
		balance.Sum += user.ChangingInBalance
		s.list[user.UserID] = balance
		userRep.Balance = s.list[user.UserID]

	} else {
		return nil, services.ChosenAccountNotFoundErr
	}
	return userRep.ToEntity(), nil
}
