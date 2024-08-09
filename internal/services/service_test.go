package services_test

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"bank-service/internal/services/mocks"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestBankService_Create1(t *testing.T) {
	type fields struct {
		bankRep *mocks.MockReposI
	}
	type args struct {
		user *services.CreateAccount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "test",
			args: args{user: &services.CreateAccount{
				UserID:  1,
				Balance: 100,
			},
			},
			want: &entity.User{ID: 1, Balance: entity.Balance{Sum: 100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				bankRep: mocks.NewMockReposI(ctrl),
			}
			got, err := f.bankRep.CreateAccount(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
