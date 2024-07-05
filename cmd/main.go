package main

import (
	"bank-service/rkg/accounts"
	"encoding/json"
	"net/http"
	"regexp"
)

var (
	AccountsRe          = regexp.MustCompile(`^/accounts/*$`)
	AccountsReWithLogin = regexp.MustCompile(`^/accounts/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func main() {

	store := accounts.NewMemStore()
	accountHandler := NewAccountHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/accounts", accountHandler)
	mux.Handle("/accounts/", accountHandler)
	http.ListenAndServe(":8080", mux)

}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

type bankStore interface {
	Add(name string, balance float64) error
	List() (map[string]float64, error)
	Show(name string) (float64, error)
}

type AccountHandler struct {
	store bankStore
}

func NewAccountHandler(b bankStore) *AccountHandler {
	return &AccountHandler{
		store: b,
	}
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && AccountsRe.MatchString(r.URL.Path):
		//Создать пользователя
		h.CreateAccount(w, r)
		return
	case r.Method == http.MethodGet && AccountsRe.MatchString(r.URL.Path):
		h.ListAccounts(w, r)
		return
	case r.Method == http.MethodGet && AccountsReWithLogin.MatchString(r.URL.Path):
		//Показать счет
		h.ShowBalance(w, r)
		return
	case r.Method == http.MethodPut && AccountsRe.MatchString(r.URL.Path):

		//Добавить деньги

		return
	case r.Method == http.MethodDelete && AccountsRe.MatchString(r.URL.Path):
		//снять Деньги
		return

	}
}

func (h *AccountHandler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	matches := AccountsReWithLogin.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}
	balance, err := h.store.Show(matches[1])

	if err != nil {
		if err == accounts.NotFoundErr {
			NotFoundHandler(w, r)
			return
		}

		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := json.Marshal(balance)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account accounts.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	if err := h.store.Add(account.Login, account.Money); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	recipesList, err := h.store.List()

	jsonBytes, err := json.Marshal(recipesList)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
