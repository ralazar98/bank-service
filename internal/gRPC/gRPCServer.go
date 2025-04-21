package servergrpc

import (
	"bank-service/internal/entity"
	"bank-service/proto"
	"context"
	"log"
	"time"
)

type BankServer struct {
	proto.UnimplementedBankServiceServer
	bankService BankServiceI
}

type BankServiceI interface {
	Create(user *entity.CreateAccount) (*entity.User, error)
	Get(user *entity.GetBalance) (*entity.User, error)
	Update(user *entity.UpdateBalance) error
}

func NewBankServer(bankService BankServiceI) *BankServer {
	return &BankServer{
		bankService: bankService,
	}
}

func (bankServer *BankServer) CreateAccount(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	var user entity.CreateAccount
	user.UserID = int(req.UserID)
	user.Balance = int(req.Balance)

	resp, err := bankServer.bankService.Create(&user)
	if err != nil {
		return nil, err
	}
	createResponse := &proto.CreateResponse{UserID: int64(resp.ID), Balance: int64(resp.Balance.Sum)}
	return createResponse, nil
}

func (bankServer *BankServer) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	var user entity.GetBalance
	user.UserID = int(req.UserID)
	resp, err := bankServer.bankService.Get(&user)
	if err != nil {

		log.Println("test")
		return nil, err
	}
	log.Println(resp.Balance)
	getBalanceResponse := &proto.GetBalanceResponse{UserID: int64(resp.ID), Balance: int64(resp.Balance.Sum)}
	return getBalanceResponse, nil
}

func (bankServer *BankServer) UpdateBalance(ctx context.Context, req *proto.UpdateBalanceRequest) (*proto.UpdateBalanceResponse, error) {
	var user entity.UpdateBalance
	user.UserID = int(req.UserID)
	user.ChangingInBalance = int(req.ChangingInBalance)
	user.Operation = req.Operation
	err := bankServer.bankService.Update(&user)
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
	resp, err := bankServer.bankService.Get(&entity.GetBalance{UserID: user.UserID})
	updateBalanceResponse := &proto.UpdateBalanceResponse{UserID: int64(resp.ID), Balance: int64(resp.Balance.Sum)}
	return updateBalanceResponse, nil
}
