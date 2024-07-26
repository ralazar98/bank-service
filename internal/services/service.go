package services

import (
	"errors"
)

var (
	NotEnoughBalanceErr      = errors.New("not enough balance")
	AccountsNotFoundErr      = errors.New("accounts not found")
	ChosenAccountNotFoundErr = errors.New("chosen account not found")
	AccountAlreadyExistsErr  = errors.New("account already exists")
)

type ReposI interface {
	CreateAccount(user *CreateAccount) *CreateAccountResponse
	GetBalance(user *GetBalance) *GetBalanceResponse
	ListAccounts() *ListAccountResponse
	UpdateBalance(user *UpdateBalance) *UpdateBalanceResponse
}

type BankService struct {
	bankRep ReposI
}

func NewBankService(bankRep ReposI) *BankService {
	return &BankService{
		bankRep: bankRep,
	}
}

func (s *BankService) CreateAccount(user *CreateAccount) (*CreateAccountResponse, error) {
	//todo AccountAlreadyExistsErr
	created := s.bankRep.CreateAccount(user)
	return created, nil
}

func (s *BankService) GetBalance(user *GetBalance) *GetBalanceResponse {
	if _, err := s.bankRep.Get(user.UserID); err != nil {
		return user.toEntity(0, ChosenAccountNotFoundErr)
	}
	return user.toEntity(s.bankRep.Get(user.UserID))

}

func (s *BankService) ListAccounts() *ListAccountResponse {
	if s.bankRep == nil {
		return &ListAccountResponse{
			List:  nil,
			Error: AccountsNotFoundErr,
		}
	}
	list, err := s.bankRep.List()
	return &ListAccountResponse{
		List:  list,
		Error: err,
	}
}

func (s *BankService) UpdateBalance(user *UpdateBalance) *UpdateBalanceResponse {
	if _, err := s.bankRep.Get(user.UserID); err != nil {
		return user.toEntity(ChosenAccountNotFoundErr)
	}
	balance, _ := s.bankRep.Get(user.UserID)
	if balance+user.ChangingInBalance < 0 {
		return user.toEntity(NotEnoughBalanceErr)
	} else {
		return user.toEntity(s.bankRep.Update(user.UserID, user.ChangingInBalance))
	}
}
