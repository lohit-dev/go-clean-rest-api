package respond

import (
	"bytes"
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
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}
