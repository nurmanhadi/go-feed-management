package response

import (
	"encoding/json"
	"feed-management/pkg/dto"
	"fmt"
	"net/http"
)

type ErrorCustom struct {
	Code    int
	Message string
}

func (e *ErrorCustom) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Message)
}
func Except(code int, message string) error {
	return &ErrorCustom{
		Code:    code,
		Message: message,
	}
}
func Success[T any](w http.ResponseWriter, code int, data T, path string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(dto.WebResponse[T]{
		Data: &data,
		Path: path,
	})
}
