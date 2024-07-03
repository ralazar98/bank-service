package main

import (
	"net/http"
	"regexp"
)

var (
	MainPage = '/'
	Action   = regexp.MustCompile(`^/action`)
)

type testHandler struct{}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &testHandler{})

	http.ListenAndServe(":8080", mux)

}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		w.Write([]byte("Это главная страница"))
	case Action.MatchString(r.URL.Path):
		w.Write([]byte("Страница действия"))

	}
}

func (t *testHandler) AddMoney(w http.ResponseWriter, r *http.Request)   {}
func (t *testHandler) GetMoney(w http.ResponseWriter, r *http.Request)   {}
func (t *testHandler) CheckMoney(w http.ResponseWriter, r *http.Request) {}
