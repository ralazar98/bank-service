package services

import (
	"bank-service/internal/entity"
	"errors"
)

var (
	NotEnoughBalanceErr      = errors.New("not enough balance")
	ChosenAccountNotFoundErr = errors.New("chosen account not found")
	AccountAlreadyExistsErr  = errors.New("account already exists")
	MinusBalanceErr          = errors.New("minus balance")
	WrongIdErr               = errors.New("wrong id")
)

type ReposI interface {
	CreateAccount(user *CreateAccount) (*entity.User, error)
	GetBalance(user *GetBalance) (*entity.User, error)
	UpdateBalance(user *UpdateBalance) (*entity.User, error)
}

type BankService struct {
	BankRep ReposI
}

func NewBankService(bankRep ReposI) *BankService {
	return &BankService{
		BankRep: bankRep,
	}
}

func (s *BankService) Create(user *CreateAccount) (*entity.User, error) {
	if user.Balance < 0 {
		return nil, MinusBalanceErr
	}
	if user.UserID < 0 {
		return nil, WrongIdErr
	}
	created, err := s.BankRep.CreateAccount(user)
	if err != nil {
		return nil, err
	}
	return created, err
}

func (s *BankService) Get(user *GetBalance) (*entity.User, error) {
	if user.UserID < 0 {
		return nil, WrongIdErr
	}
	gotBalance, err := s.BankRep.GetBalance(user)
	return gotBalance, err

}

func (s *BankService) Update(user *UpdateBalance) (*entity.User, error) {

	updatedBalance, err := s.BankRep.UpdateBalance(user)
	return updatedBalance, err

}
