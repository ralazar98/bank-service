package services

type CreateAccount struct {
	UserID  int     `json:"userID"`
	Balance float64 `json:"balance"`
}

type ListAccounts struct {
	ListAccount map[int]float64 `json:"listAccount"`
}

type GetBalance struct {
	UserID int `json:"userID"`
}

type UpdateBalance struct {
	UserID            int     `json:"userID"`
	Operation         string  `json:"operation"`
	ChangingInBalance float64 `json:"changing_in_balance"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
	Error   error   `json:"error,omitempty"`
}

func toEntity() *GetBalanceResponse {
	return &GetBalanceResponse{
		Balance: balance,
		Error:   err,
	}

}

func (c *UpdateBalance) toEntity() {}
