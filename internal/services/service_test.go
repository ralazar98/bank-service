package services_test

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"bank-service/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestBankService_Create(t *testing.T) {
	type fields struct {
		BankRep *mock_services.MockReposI
	}
	type args struct {
		user *services.CreateAccount
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "base_test",
			args: args{user: &services.CreateAccount{UserID: 1, Balance: 100}},
			want: &entity.User{ID: 1, Balance: entity.Balance{Sum: 100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bankRepMock := mock_services.NewMockReposI(ctrl)

			got, err := bankRepMock.CreateAccount(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBankService_Get(t *testing.T) {
	type fields struct {
		BankRep *mock_services.MockReposI
	}
	type args struct {
		user *services.GetBalance
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{BankRep: mock_services.NewMockReposI(ctrl)}

			s := NewBankService(f.BankRep)
			got, err := s.Get(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
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
	tests := []struct {
		name string

		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{BankRep: mock_services.NewMockReposI(ctrl)}

			s := &services.BankService{
				BankRep: f.BankRep,
			}
			got, err := s.Update(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBankService(t *testing.T) {
	type args struct {
		bankRep *mock_services.MockReposI
	}
	tests := []struct {
		name string
		args args
		want *services.BankService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBankService(tt.args.bankRep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBankService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewBankService(rep *mock_services.MockReposI) *services.BankService {
	return &services.BankService{
		BankRep: rep,
	}
}
