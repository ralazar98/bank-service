package entity

type User struct {
	ID      int     `json:"id" `
	Balance Balance `json:"balance" `
}

type Balance struct {
	Sum int `json:"sum" `
}

// TODO: поля в json обычно отдаются в snake_case, а не camelCase
type UpdateBalance struct {
	UserID            int    `json:"userID" validate:"gte=1"`
	Operation         string `json:"operation" validate:"required" `
	ChangingInBalance int    `json:"changingInBalance" validate:"gte=1"`
}

type CreateAccount struct {
	UserID  int `json:"userID" validate:"gte=1"`
	Balance int `json:"balance" validate:"gte=0"`
}

type GetBalance struct {
	UserID int `json:"userID" validate:"gte=1"`
}
