package http

type CreateAccountRequest struct {
	UserID  int     `json:"userID"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	UserID            int     `json:"userID"`
	ChangingInBalance float64 `json:"changing_in_balance"`
	Operation         string  `json:"operation"`
}

type ShowBalance struct {
	UserID int `json:"userID"`
}

type UpdateBalanceRequest struct {
	UserID            int     `json:"userID"`
	Operation         string  `json:"operation"`
	ChangingInBalance float64 `json:"changing_in_balance"`
}
