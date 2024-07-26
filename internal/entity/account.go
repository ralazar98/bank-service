package entity

type User struct {
	ID      int
	Balance Balance
}

type Balance struct {
	Sum int
}
