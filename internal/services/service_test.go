package services_test

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"bank-service/mocks"
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
			name: "base_test",
			args: args{CreateAccountUser(1, 100)},
			want: CreateEntityUser(1, 100),
			prepare: func() {
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(1, 100)).
					Return(CreateEntityUser(1, 100), nil)
			},
			wantErr: nil,
		},
		{
			name: "AccountAlreadyExists_test",
			args: args{CreateAccountUser(1, 100)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().CreateAccount(CreateAccountUser(1, 100)).
					Return(nil, services.AccountAlreadyExistsErr)
			},
			wantErr: services.AccountAlreadyExistsErr,
		},
		/*		{
					name: "minus_balance_test",
					args: args{CreateAccountUser(2, -100)},
					want: nil,
					prepare: func() {
						bankRepMock.EXPECT().CreateAccount(CreateAccountUser(2, -100)).
							Return(nil, services.MinusBalanceErr)
					},
					wantErr: services.MinusBalanceErr,
				},
				{
					name: "wrong_ID_test",
					args: args{CreateAccountUser(-1, 100)},
					want: nil,
					prepare: func() {
						bankRepMock.EXPECT().GetBalance(CreateAccountUser(-1, 100)).Return(nil, services.WrongIdErr)
					},
					wantErr: services.WrongIdErr,
				},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := service.Create(tt.args.user)
			if err != nil {
				assert.Errorf(t, err, "error")
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
			name: "base_test",
			args: args{GetBalanceUser(1)},
			want: CreateEntityUser(1, 100),
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(1)).Return(CreateEntityUser(1, 100), nil)
			},
			wantErr: nil,
		},
		{
			name: "not_exists_test",
			args: args{GetBalanceUser(6)},
			want: nil,
			prepare: func() {
				bankRepMock.EXPECT().GetBalance(GetBalanceUser(6)).Return(nil, services.ChosenAccountNotFoundErr)
			},
			wantErr: services.ChosenAccountNotFoundErr,
		},
		/*		{
				name: "wrong_ID_test",
				args: args{GetBalanceUser(-1)},
				want: nil,
				prepare: func() {
					bankRepMock.EXPECT().GetBalance(GetBalanceUser(-1)).Return(nil, services.WrongIdErr)
				},
				wantErr: services.WrongIdErr,
			},*/
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
		BankRep *mock_services.MockReposI
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
			name: "add_money_test",
			args: args{user: UpdateBalanceUser(1, 50)},
			want: CreateEntityUser(1, 150),
			prepare: func() {
				bankRepMock.EXPECT().UpdateBalance(UpdateBalanceUser(1, 50)).
					Return(CreateEntityUser(1, 150), nil)
			},
			wantErr: nil,
		},
		{
			name: "take_money_test",
			args: args{user: UpdateBalanceUser(1, -50)},
			want: CreateEntityUser(1, 50),
			prepare: func() {
				bankRepMock.EXPECT().UpdateBalance(UpdateBalanceUser(1, -50)).
					Return(CreateEntityUser(1, 50), nil)
			},
			wantErr: nil,
		},
		{
			name: "not_enough_money_test",
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

func NewBankService(rep *mock_services.MockReposI) *services.BankService {
	return &services.BankService{
		BankRep: rep,
	}
}
