package errs

import "net/http"

type Error struct {
    Code int `json:"code,omitempty"`
    Message string `json:"message"`
}

func NotFoundError(msg string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ServerError(msg string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func UnAuthorizedError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}
