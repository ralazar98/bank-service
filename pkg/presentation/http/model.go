package http

import (
	"bank-service/internal/services"
	"bank-service/internal/storage"
)

type CreateAccount struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	ID                int               `json:"id"`
	ChangingInBalance float64           `json:"changing_in_balance"`
	Operation         storage.Operation `json:"operation"`
}

type ShowBalance struct {
	ID int `json:"id"`
}

func (c *CreateAccount) ToEntity() *services.CreateAcc {
	return &services.CreateAcc{
		ID:      c.ID,
		Balance: c.Balance,
	}

}
