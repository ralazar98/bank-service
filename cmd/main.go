package main

import (
	"bank-service/config"
	"bank-service/internal/storage"
	"net/http"
)

// Объявление переменных

func main() {

	store := storage.New()

	cfg := config.New()
	mux := http.NewServeMux()

	if err := http.ListenAndServe(cfg.Server.Port, mux); err != nil {
		panic(err)
	}

}

// Обработчики ошибок
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

// Интерфейс хранения
type bankStore interface {
	Add(name string, balance float64) error
	List() (map[string]float64, error)
	Show(name string) (float64, error)
	ChangeBalance(name string, operation string, changingInBalance float64) error
}
