package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultError struct {
	Message string `json:"message"`
}

func DecodeRequest[T any](r *http.Request) (T, error) {
	var req T
	return req, json.NewDecoder(r.Body).Decode(&req)
}

func WriteErrorResponse(w http.ResponseWriter, code int, format string, a ...any) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(DefaultError{
		Message: fmt.Sprintf(format, a...),
	})
}
