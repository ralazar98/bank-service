package postgresql

import (
	"bank-service/internal/entity"
)

func ToEntity(id int, balance int) *entity.User {
	return &entity.User{
		ID: id,
		Balance: entity.Balance{
			Sum: balance,
		},
	}
}
