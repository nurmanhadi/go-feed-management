package middleware

import (
	"feed-management/pkg/dto"
	"feed-management/pkg/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				w.Header().Set("Content-Type", "application/json")
				switch err := rec.(type) {
				case validator.ValidationErrors:
					var data []string
					for _, x := range err {
						value := fmt.Sprintf("field %s is %s %s", x.Field(), x.Tag(), x.Param())
						data = append(data, value)
					}
					msg := strings.Join(data, ", ")
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(dto.WebResponse[string]{
						Error: &msg,
						Path:  r.URL.Path,
					})
					return
				case *response.ErrorCustom:
					w.WriteHeader(err.Code)
					json.NewEncoder(w).Encode(dto.WebResponse[string]{
						Error: &err.Message,
						Path:  r.URL.Path,
					})
					return
				default:
					msg := fmt.Sprintf("%v", err)
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(dto.WebResponse[string]{
						Error: &msg,
						Path:  r.URL.Path,
					})
					return
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
