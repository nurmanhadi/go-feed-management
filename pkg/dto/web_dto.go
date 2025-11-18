package dto

type WebResponse[T any] struct {
	Data  *T      `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
	Path  string  `json:"path"`
}
