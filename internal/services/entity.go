package services

type CreateAccount struct {
	UserID  int `json:"userID"`
	Balance int `json:"balance"`
}

type GetBalance struct {
	UserID int `json:"userID"`
}

type UpdateBalance struct {
	UserID            int    `json:"userID"`
	Operation         string `json:"operation"`
	ChangingInBalance int    `json:"changing_in_balance"`
}
