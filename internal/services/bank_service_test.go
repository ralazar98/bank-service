package services_test

import (
	"bank-service/internal/entity"
	"bank-service/internal/mocks"
	"bank-service/internal/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func CreateAccountUser(userId int, balance int) *services.CreateAccount {
	return &services.CreateAccount{
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

func GetBalanceUser(userId int) *services.GetBalance {
	return &services.GetBalance{
		UserID: userId,
	}
}

func UpdateBalanceUser(userId int, changingInBalance int) *services.UpdateBalance {
	return &services.UpdateBalance{
		UserID:            userId,
		ChangingInBalance: changingInBalance,
	}
}

func TestBankService_Create(t *testing.T) {
	type args struct {
		user *services.CreateAccount
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mock_services.mock_services.NewMockReposI(ctrl)
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
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(1, 100)).
					Return(nil, services.AccountAlreadyExistsErr)
			},
			wantErr: services.AccountAlreadyExistsErr,
		},
		{
			name: "CreateAccount Failed, minus balance",
			args: args{CreateAccountUser(10, -100)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(2, -100)).
					Return(nil, services.MinusBalanceErr)
			},
			wantErr: services.MinusBalanceErr,
		},
		{
			name: "CreateAccount Failed, wrong ID",
			args: args{CreateAccountUser(-1, 100)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(-1, 100)).Return(nil, services.WrongIdErr)
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
		user *services.GetBalance
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mock_services.NewMockReposI(ctrl)
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
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(-1)).Return(nil, services.WrongIdErr)
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
	type fields struct {
		BankRep *mock_services.mock_services
	}
	type args struct {
		user *services.UpdateBalance
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bankRepMock := mock_services.NewMockReposI(ctrl)
	service := NewBankService(bankRepMock)

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		prepare func()
		wantErr error
	}{
		{
			name: "Update(add) balance Success",
			args: args{user: UpdateBalanceUser(1, 50)},
			want: CreateEntityUser(1, 150),
			prepare: func() {
				bankRepMock.EXPECT().UpdateBalance(UpdateBalanceUser(1, 50)).
					Return(CreateEntityUser(1, 150), nil)
			},
			wantErr: nil,
		},
		{
			name: "Update(take) balance Success",
			args: args{user: UpdateBalanceUser(1, -50)},
			want: CreateEntityUser(1, 50),
			prepare: func() {
				bankRepMock.EXPECT().UpdateBalance(UpdateBalanceUser(1, -50)).
					Return(CreateEntityUser(1, 50), nil)
			},
			wantErr: nil,
		},
		{
			name: "Update take(failed) failed,not enough balance",
			args: args{user: UpdateBalanceUser(1, -150)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().UpdateBalance(UpdateBalanceUser(1, -150)).
					Return(nil, services.NotEnoughBalanceErr)
			},
			wantErr: services.NotEnoughBalanceErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := service.Update(tt.args.user)
			if err != nil {
				assert.Errorf(t, err, "error")
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func NewBankService(rep *mock_services.mock_services) *services.BankService {
	return &services.BankService{
		BankRep: rep,
	}
}
