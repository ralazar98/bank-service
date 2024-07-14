package services

type CreateAcc struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	ID                int     `json:"id"`
	ChangingInBalance float64 `json:"changing_in_balance"`
	Operation         int     `json:"operation"`
}
