package accounts

type Account struct {
	ID                string  `json:"id"`
	Balance           float64 `json:"balance"`
	Operation         string  `json:"operation"`
	ChangingInBalance float64 `json:"changing_in_balance"`
}
