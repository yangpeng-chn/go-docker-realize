package app

import (
	"log"
	"net/http"

	"github.com/yangpeng-chn/go-docker-realize/app/controller"

	"github.com/gorilla/mux"
)

var apiKeys = map[string]uint32{
	"secure-api-key-1": 1,
	"secure-api-key-2": 2,
	"secure-api-key-3": 3,
}

// NewMux create the handler.
func NewMux() http.Handler {
	r := mux.NewRouter()

	r.Use(authenticate)
	r.Use(requestLogger)

	transaction := controller.TransactionController{}
	r.HandleFunc("/transactions", transaction.List).Methods(http.MethodGet)
	r.HandleFunc("/transactions/{id:[0-9]+}", transaction.Show).Methods(http.MethodGet)
	r.HandleFunc("/transactions", transaction.Create).Methods(http.MethodPost)
	r.HandleFunc("/transactions", transaction.Delete).Methods(http.MethodDelete)

	return r
}

func authenticate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("apikey")
		if _, ok := apiKeys[apiKey]; !ok {
			http.Error(w, "invalid apikey", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", r)
		handler.ServeHTTP(w, r)
	})
}
