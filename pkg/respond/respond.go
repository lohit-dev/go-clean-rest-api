package respond

import (
	"encoding/json"
	"net/http"
)

type Status string

const (
	StatusOk    Status = "Ok"
	StatusError Status = "Error"
)

type ApiResponse[T any] struct {
	Status Status `json:"status"`
	Result T      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func Ok[T any](w http.ResponseWriter, status int, data T) {
	write(w, status, ApiResponse[T]{Status: StatusOk, Result: data})
}

func ErrorMessage(w http.ResponseWriter, status int, msg string) {
	write(w, status, ApiResponse[any]{Status: StatusError, Error: msg})
}

func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
