package entity

type User struct {
	ID      int     `json:"id"`
	Balance Balance `json:"balance"`
}

type Balance struct {
	Sum int `json:"sum"`
}

type UpdateBalance struct {
	UserID            int    `json:"userID"`
	Operation         string `json:"operation"`
	ChangingInBalance int    `json:"changingInBalance"`
}

type CreateAccount struct {
	UserID  int `json:"userID"`
	Balance int `json:"balance"`
}

type GetBalance struct {
	UserID int `json:"userID"`
}
