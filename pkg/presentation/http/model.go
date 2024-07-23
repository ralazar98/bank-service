package http

import (
	"bank-service/internal/services"
)

type CreateAccountRequest struct {
	UserID  int     `json:"userID"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	ID                int     `json:"id"`
	ChangingInBalance float64 `json:"changing_in_balance"`
	Operation         string  `json:"operation"`
}

type ShowBalance struct {
	ID int `json:"id"`
}

func (c *CreateAccountRequest) ToEntity() *services.CreateAcc {
	return &services.CreateAcc{
		ID:      c.UserID,
		Balance: c.Balance,
	}
}
