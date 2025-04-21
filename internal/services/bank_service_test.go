package services_test

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"bank-service/mocks"
	_ "bank-service/mocks"
	mock_rabbit "bank-service/mocks/rabbitmq"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CreateAccountUser(userId int, balance int) *entity.CreateAccount {
	return &entity.CreateAccount{
		UserID:  userId,
		Balance: balance,
	}
}

func CreateEntityUser(userId int, balance int) *entity.User {
	return &entity.User{
		ID: userId,
		Balance: entity.Balance{
			Sum: balance,
		},
	}
}

func GetBalanceUser(userId int) *entity.GetBalance {
	return &entity.GetBalance{
		UserID: userId,
	}
}

func UpdateBalanceUser(userId int, changingInBalance int) *entity.UpdateBalance {
	return &entity.UpdateBalance{
		UserID:            userId,
		ChangingInBalance: changingInBalance,
	}
}

func TestBankService_Create(t *testing.T) {
	type args struct {
		user *entity.CreateAccount
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mocks.NewMockReposI(ctrl)
	service := NewBankService(bankRepMock)

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		prepare func()
		wantErr error
	}{
		{
			name: "CreateAccount Success",
			args: args{CreateAccountUser(1, 100)},
			want: CreateEntityUser(1, 100),
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(1)).
					Return(nil, services.ChosenAccountNotFoundErr)
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(1, 100)).
					Return(CreateEntityUser(1, 100), nil)

			},
			wantErr: nil,
		},
		{
			name: "CreateAccount Failed, already exist",
			args: args{CreateAccountUser(1, 100)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(1)).
					Return(CreateEntityUser(1, 100), nil)

			},
			wantErr: services.AccountAlreadyExistsErr,
		},
		{
			name: "CreateAccount Failed, minus balance",
			args: args{CreateAccountUser(10, -100)},
			want: nil,
			prepare: func() {

			},
			wantErr: services.MinusBalanceErr,
		},
		{
			name: "CreateAccount Failed, wrong ID",
			args: args{CreateAccountUser(-1, 100)},
			want: nil,
			prepare: func() {

			},
			wantErr: services.WrongIdErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := service.Create(tt.args.user)
			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBankService_Get(t *testing.T) {
	type args struct {
		user *entity.GetBalance
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mocks.NewMockReposI(ctrl)
	service := NewBankService(bankRepMock)

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		prepare func()
		wantErr error
	}{
		{
			name: "Get balance Success",
			args: args{GetBalanceUser(1)},
			want: CreateEntityUser(1, 100),
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(1)).Return(CreateEntityUser(1, 100), nil)
			},
			wantErr: nil,
		},
		{
			name: "Get balance Failed, not found",
			args: args{GetBalanceUser(6)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(6)).Return(nil, services.ChosenAccountNotFoundErr)
			},
			wantErr: services.ChosenAccountNotFoundErr,
		},
		{
			name: "Get balance Failed, wrong ID",
			args: args{GetBalanceUser(-1)},
			want: nil,
			prepare: func() {

			},
			wantErr: services.WrongIdErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := service.Get(tt.args.user)
			if err != nil {
				assert.Errorf(t, err, "error")
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBankService_Update(t *testing.T) {

	type args struct {
		user *entity.UpdateBalance
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mocks.NewMockReposI(ctrl)
	service := NewBankService(bankRepMock)
	rabbit := mock_rabbit.NewMockRabbitI(ctrl)

	tests := []struct {
		name    string
		args    args
		prepare func()
		wantErr error
	}{
		{
			name: "Update(add) balance Success",
			args: args{user: UpdateBalanceUser(1, 50)},
			prepare: func() {
				service.Get(GetBalanceUser(1))
				//bankRepMock.EXPECT().GetBalance(GetBalanceUser(1)).Return(CreateEntityUser(1, 50))
				rabbit.EXPECT().SendToPaymentService(UpdateBalanceUser(1, 50)).
					Return(nil)
				bankRepMock.EXPECT().Update(UpdateBalanceUser(1, 50)).
					Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "Update(take) balance Success",
			args: args{user: UpdateBalanceUser(1, -50)},
			prepare: func() {
				bankRepMock.EXPECT().Update(UpdateBalanceUser(1, -50)).
					Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "Update take(failed) failed,not enough balance",
			args: args{user: UpdateBalanceUser(1, -150)},
			prepare: func() {
				bankRepMock.EXPECT().Update(UpdateBalanceUser(1, -150)).
					Return(services.NotEnoughBalanceErr)
			},
			wantErr: services.NotEnoughBalanceErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			err := service.Update(tt.args.user)
			assert.Errorf(t, err, "error")
		})
	}
}

func NewBankService(rep *mocks.MockReposI) *services.BankService {
	return &services.BankService{
		BankRep: rep,
	}
}
