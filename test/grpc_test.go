package test

import (
	"bank-service/internal/services"
	"bank-service/mocks"
	"bank-service/proto"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CreateAccountRequest(userID int64, balance int64) *proto.CreateRequest {
	return &proto.CreateRequest{
		UserID:  userID,
		Balance: balance,
	}
}
func CreateAccountResponse(userID int64, balance int64) *proto.CreateResponse {
	return &proto.CreateResponse{
		UserID:  userID,
		Balance: balance,
	}
}

func GetBalanceRequest(userID int64) *proto.GetBalanceRequest {
	return &proto.GetBalanceRequest{
		UserID: userID,
	}
}

func GetBalanceResponse(userID int64, balance int64) *proto.GetBalanceResponse {
	return &proto.GetBalanceResponse{
		UserID:  userID,
		Balance: balance,
	}
}

func UpdateBalanceRequest(userID int64, changingInBalance int64, operation string) *proto.UpdateBalanceRequest {
	return &proto.UpdateBalanceRequest{
		UserID:            userID,
		Operation:         operation,
		ChangingInBalance: changingInBalance,
	}
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		req *proto.CreateRequest
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBankService := mocks.NewMockBankServiceClient(ctrl)

	tests := []struct {
		name    string
		args    args
		want    *proto.CreateResponse
		wantErr error
		prepare func()
	}{
		{
			name: "CreateAccount:Success",
			args: args{
				req: CreateAccountRequest(1, 100),
			},
			want:    CreateAccountResponse(1, 100),
			wantErr: nil,
			prepare: func() {
				mockBankService.EXPECT().
					CreateAccount(gomock.Any(), CreateAccountRequest(1, 100)).
					Return(CreateAccountResponse(1, 100), nil)
			},
		},
		{
			name: "CreateAccount:Fail,AccountAlreadyExists",
			args: args{
				req: CreateAccountRequest(1, 1000),
			},
			want:    nil,
			wantErr: services.AccountAlreadyExistsErr,
			prepare: func() {
				mockBankService.EXPECT().
					CreateAccount(gomock.Any(), CreateAccountRequest(1, 1000)).
					Return(nil, services.AccountAlreadyExistsErr)
			},
		},
		{
			name: "CreateAccount:Fail,MinusBalanceErr",
			args: args{
				req: CreateAccountRequest(1, -1000),
			},
			want:    nil,
			wantErr: services.MinusBalanceErr,
			prepare: func() {
				mockBankService.EXPECT().CreateAccount(gomock.Any(), CreateAccountRequest(1, -1000)).
					Return(nil, services.MinusBalanceErr)
			},
		},
		{
			name: "CreateAccount:Fail,WrongIDErr",
			args: args{
				req: CreateAccountRequest(-1, 1000),
			},
			want:    nil,
			wantErr: services.WrongIdErr,
			prepare: func() {
				mockBankService.EXPECT().CreateAccount(gomock.Any(), CreateAccountRequest(-1, 1000)).
					Return(nil, services.WrongIdErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := mockBankService.CreateAccount(context.Background(), tt.args.req)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, res)
		})
	}

}

func TestGetBalance(t *testing.T) {
	type args struct {
		req *proto.GetBalanceRequest
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBankService := mocks.NewMockBankServiceClient(ctrl)
	tests := []struct {
		name    string
		args    args
		want    *proto.GetBalanceResponse
		wantErr error
		prepare func()
	}{
		{
			name: "GetBalance:Success",
			args: args{
				GetBalanceRequest(1),
			},
			want:    GetBalanceResponse(1, 100),
			wantErr: nil,
			prepare: func() {
				mockBankService.EXPECT().GetBalance(gomock.Any(), GetBalanceRequest(1)).
					Return(GetBalanceResponse(1, 100), nil)
			},
		},
		{
			name: "GetBalance:Fail,AccountNotExists",
			args: args{
				GetBalanceRequest(100),
			},
			want:    nil,
			wantErr: services.ChosenAccountNotFoundErr,
			prepare: func() {
				mockBankService.EXPECT().GetBalance(gomock.Any(), GetBalanceRequest(100)).
					Return(nil, services.ChosenAccountNotFoundErr)
			},
		},
		{
			name: "GetBalance:Fail,WrongIDErr",
			args: args{
				GetBalanceRequest(-100),
			},
			want:    nil,
			wantErr: services.WrongIdErr,
			prepare: func() {
				mockBankService.EXPECT().GetBalance(gomock.Any(), GetBalanceRequest(-100)).
					Return(nil, services.WrongIdErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := mockBankService.GetBalance(context.Background(), tt.args.req)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestUpdateBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBankService := mocks.NewMockBankServiceClient(ctrl)
	//mockRabbit := mock_rabbit.NewMockRabbitI(ctrl)
	type args struct {
		req *proto.UpdateBalanceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		prepare func()
	}{
		{
			name: "UpdateBalance:Success",
			args: args{
				UpdateBalanceRequest(1, 100, "пополнить"),
			},
			wantErr: nil,
			prepare: func() {
				//mockRabbit.EXPECT().SendToPaymentService(UpdateBalanceRequest(1, 100, "пополнить"))
				mockBankService.EXPECT().UpdateBalance(gomock.Any(), UpdateBalanceRequest(1, 100, "пополнить")).Return(gomock.Any(), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			_, err := mockBankService.UpdateBalance(context.Background(), tt.args.req)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}

}
