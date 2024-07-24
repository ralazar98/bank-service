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
	Create(userID int, balance float64) error
	Get(userID int) (float64, error)
	List() (map[int]float64, error)
	Update(userID int, changingInBalance float64) error
}

type BankService struct {
	bankRep *memory_cache.BankStorage
}

func NewBankService(bankRep *memory_cache.BankStorage) *BankService {
	return &BankService{
		bankRep,
	}
}

func (s *BankService) Create(userID int, balance float64) error {
	if _, err := s.bankRep.Get(userID); err == nil {
		return AccountAlreadyExistsErr
	}
	return s.bankRep.Create(userID, balance)
}
func (s *BankService) Get(userID int) (float64, error) {
	if _, err := s.bankRep.Get(userID); err != nil {
		return 0, ChosenAccountNotFoundErr
	}
	return s.bankRep.Get(userID)
}

func (s *BankService) List() (map[int]float64, error) {
	if s.bankRep == nil {
		return nil, AccountsNotFoundErr
	}
	return s.bankRep.List()
}

func (s *BankService) Update(userID int, changingInBalance float64) error {
	if _, err := s.bankRep.Get(userID); err != nil {
		return ChosenAccountNotFoundErr
	}
	balance, _ := s.bankRep.Get(userID)
	if balance+changingInBalance < 0 {
		return NotEnoughBalanceErr
	} else {
		return s.bankRep.Update(userID, changingInBalance)
	}
}
