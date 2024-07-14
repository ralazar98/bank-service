package http

import "bank-service/internal/services"

type CreateAccount struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	ID                int    `json:"id"`
	ChangingInBalance int    `json:"changing_in_balance"`
	Operation         string `json:"operation"`
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
