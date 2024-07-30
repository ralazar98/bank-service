package memory_cache

import "bank-service/internal/entity"

type User struct {
	ID      int
	Balance Balance
}

type Balance struct {
	Sum int
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID: u.ID,
		Balance: entity.Balance{
			Sum: u.Balance.Sum,
		},
	}
}
