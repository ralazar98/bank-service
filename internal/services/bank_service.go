package services

import (
	"bank-service/configs"
	"bank-service/internal/entity"
	"bank-service/internal/rabbit"
	"errors"
)

var (
	NotEnoughBalanceErr      = errors.New("not enough balance")
	ChosenAccountNotFoundErr = errors.New("chosen account not found")
	AccountAlreadyExistsErr  = errors.New("account already exists")
	MinusBalanceErr          = errors.New("minus balance")
	WrongIdErr               = errors.New("wrong id")
	InvalidOperationErr      = errors.New("invalid operation")
)

type ReposI interface {
	CreateAccount(user *entity.CreateAccount) (*entity.User, error)
	GetBalance(user *entity.GetBalance) (*entity.User, error)
	Update(user *entity.UpdateBalance) error
}

type BankService struct {
	BankRep ReposI
	cfg     *configs.Config
	rabbit  *rabbit.Rabbit
}

func NewBankService(bankRep ReposI, cfg *configs.Config, newRabbit *rabbit.Rabbit) *BankService {
	return &BankService{
		BankRep: bankRep,
		cfg:     cfg,
		rabbit:  newRabbit,
	}
}

func (s *BankService) Create(user *entity.CreateAccount) (*entity.User, error) {
	if user.Balance < 0 {
		return nil, MinusBalanceErr
	}

	if user.UserID < 0 {
		return nil, WrongIdErr
	}

	_, err := s.Get(&entity.GetBalance{UserID: user.UserID})
	if err == nil {
		return nil, AccountAlreadyExistsErr
	}

	created, err := s.BankRep.CreateAccount(user)
	if err != nil {
		return nil, err
	}

	return created, err
}

func (s *BankService) Get(user *entity.GetBalance) (*entity.User, error) {
	if user.UserID < 0 {
		return nil, WrongIdErr
	}

	gotBalance, err := s.BankRep.GetBalance(user)
	if err != nil {
		return nil, ChosenAccountNotFoundErr
	}

	return gotBalance, err

}

func (s *BankService) Update(user *entity.UpdateBalance) error {
	if user.UserID < 0 {
		return WrongIdErr
	}
	_, err := s.Get(&entity.GetBalance{UserID: user.UserID})
	if err != nil {
		return ChosenAccountNotFoundErr
	}

	err = s.rabbit.SendToPaymentService(user)
	if err != nil {
		return err
	}

	return nil
}
