package handlers

import (
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountHandler_CreateAccount(t *testing.T) {
	bankRep := memory_cache.New()
	service := services.NewBankService(bankRep)
	handler := NewAccountHandler(service)

	testTable := []struct {
		name     string
		body     io.Reader
		wantBody string
	}{
		{
			name:     "base",
			body:     bytes.NewReader([]byte(`{"userID":"1","balance":"100"}`)),
			wantBody: `{"userID":"1","balance":"100"}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodPost, "/create", tt.body)

			if err != nil {
				log.Fatal(err)
			}
			rr := httptest.NewRecorder()
			log.Println("Test1")
			handler.CreateAccount(rr, req)
			log.Println("Test2")
			res := rr.Result()
			defer res.Body.Close()
			body, errBody := io.ReadAll(res.Body)
			if errBody != nil {
				assert.Fail(t, err.Error())
			}
			assert.Equal(t, string(body), tt.wantBody)

		})
	}
}
