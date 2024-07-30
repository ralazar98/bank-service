package services

import (
	"bank-service/internal/entity"
	"errors"
)

var (
	NotEnoughBalanceErr      = errors.New("not enough balance")
	ChosenAccountNotFoundErr = errors.New("chosen account not found")
	AccountAlreadyExistsErr  = errors.New("account already exists")
)

type ReposI interface {
	CreateAccount(user *CreateAccount) (*entity.User, error)
	GetBalance(user *GetBalance) (*entity.User, error)
	UpdateBalance(user *UpdateBalance) (*entity.User, error)
}

type BankService struct {
	bankRep ReposI
}

func NewBankService(bankRep ReposI) *BankService {
	return &BankService{
		bankRep: bankRep,
	}
}

func (s *BankService) Create(user *CreateAccount) (*entity.User, error) {

	created, err := s.bankRep.CreateAccount(user)
	return created, err
}

func (s *BankService) Get(user *GetBalance) (*entity.User, error) {
	gotBalance, err := s.bankRep.GetBalance(user)
	return gotBalance, err

}

func (s *BankService) Update(user *UpdateBalance) (*entity.User, error) {

	updatedBalance, err := s.bankRep.UpdateBalance(user)
	return updatedBalance, err

}
