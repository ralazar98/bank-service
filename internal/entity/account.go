package entity

type Accounts struct {
	Users []User `json:"users"`
}

type User struct {
	ID      int     `json:"id"`
	Balance Balance `json:"balance"`
}

type Balance struct {
	Sum int `json:"sum"`
}
