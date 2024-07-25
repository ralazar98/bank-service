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

type CreateAccountResponse struct {
	Error error `json:"error"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
	Error   error   `json:"error"`
}

type ListAccountResponse struct {
	List  map[int]float64 `json:"list"`
	Error error           `json:"error"`
}

type UpdateBalanceResponse struct {
	Error error `json:"error"`
}

func (c *CreateAccount) toEntity(err error) *CreateAccountResponse {
	return &CreateAccountResponse{
		Error: err,
	}
}

func (g *GetBalance) toEntity(bal float64, err error) *GetBalanceResponse {
	return &GetBalanceResponse{
		Balance: bal,
		Error:   err,
	}

}

func (c *UpdateBalance) toEntity(err error) *UpdateBalanceResponse {
	return &UpdateBalanceResponse{
		Error: err,
	}
}
