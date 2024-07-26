package memory_cache

import "bank-service/internal/entity"

type balance struct {
	b float64
}

func (b *balance) ToEntity(userID int) *entity.User {
	return &entity.User{
		ID:      userID,
		Balance: b.b,
	}
}
