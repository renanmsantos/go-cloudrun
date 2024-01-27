package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmoreirasan/go-cloudrun/internal/usecases"
)

func Get() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("cep")
		output, err := usecases.Execute(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
