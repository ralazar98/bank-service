package http

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

type UpdateBalanceRequest struct {
	UserID            int     `json:"userID"`
	Operation         string  `json:"operation"`
	ChangingInBalance float64 `json:"changing_in_balance"`
}
