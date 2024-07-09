package http

import "bank-service/internal/services"

type TakeBalance struct {
	ID int `json:"id"`
}

func (t *TakeBalance) ToEntity() *services.TakeBalance {
	return &services.TakeBalance{
		ID: t.ID,
	}
}

type AddBalance struct {
}
