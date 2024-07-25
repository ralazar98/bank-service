package services

import (
	"bank-service/pkg/infrastructure/memory_cache"
	"errors"
)

var (
	NotEnoughBalanceErr      = errors.New("not enough balance")
	AccountsNotFoundErr      = errors.New("accounts not found")
	ChosenAccountNotFoundErr = errors.New("chosen account not found")
	AccountAlreadyExistsErr  = errors.New("account already exists")
)

type BankServiceI interface {
	CreateAccount(user *CreateAccount) error
	GetBalance(user *GetBalance) (float64, error)
	ListAccounts() (map[int]float64, error)
	UpdateBalance(user *UpdateBalance) error
}

type BankService struct {
	bankRep *memory_cache.BankStorage
}

func NewBankService(bankRep *memory_cache.BankStorage) *BankService {
	return &BankService{
		bankRep: bankRep,
	}
}

func (s *BankService) CreateAccount(user *CreateAccount) error {
	//todo AccountAlreadyExistsErr
	return s.bankRep.Create(user.UserID, user.Balance)
}
func (s *BankService) GetBalance(user *GetBalance) (GetBalance.) {
	if _, err := s.bankRep.Get(user.UserID); err != nil {
		return 0, ChosenAccountNotFoundErr
	}
	return s.bankRep.Get(user.UserID)
}

func (s *BankService) ListAccounts() (map[int]float64, error) {
	if s.bankRep == nil {
		return nil, AccountsNotFoundErr
	}
	return s.bankRep.List()
}

func (s *BankService) UpdateBalance(user *UpdateBalance) error {
	if _, err := s.bankRep.Get(user.UserID); err != nil {
		return ChosenAccountNotFoundErr
	}
	balance, _ := s.bankRep.Get(user.UserID)
	if balance+user.ChangingInBalance < 0 {
		return NotEnoughBalanceErr
	} else {
		return s.bankRep.Update(user.UserID, balance+user.ChangingInBalance)
	}
}
