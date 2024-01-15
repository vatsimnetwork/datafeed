package server

import (
	"encoding/json"
	"net/http"
)

func stringHandler(content *string, mimeType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeType)
		w.Write([]byte(*content))
	})
}

func jsonMarshalHandler(content interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(content)
	})
}
