package utils

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"status"`
}
