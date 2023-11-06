package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type ApiError interface {
	WithMessage(string) ApiError
	WithCode(int) ApiError
	Error() error

	getCode() string
	build() *EventResponse
}

type apiErrorBuilder struct {
	status  int
	code    int
	message string
}

func NewApiError(status int) ApiError {
	return &apiErrorBuilder{
		status: status,
	}
}

func (b *apiErrorBuilder) WithStatus(code int) ApiError {
	b.status = code

	return b
}

func (b *apiErrorBuilder) WithMessage(message string) ApiError {
	b.message = message

	return b
}

func (b *apiErrorBuilder) WithCode(code int) ApiError {
	b.code = code

	return b
}

func (b *apiErrorBuilder) Error() error {
	return errors.New(b.message)
}

func (b *apiErrorBuilder) getCode() string {
	return fmt.Sprintf("E%d", b.code)
}

func (b *apiErrorBuilder) build() *EventResponse {
	log.Printf("error:\"%s\", status:%d, code:%s", b.message, b.status, b.getCode())

	if b.status == http.StatusInternalServerError {
		b.message = http.StatusText(b.status)
	}

	return NewResponse(b.status).
		WithBody(b.message).
		WithHeader("code", b.getCode()).
		build()
}

func ToApiError(err error) ApiError {
	return NewApiError(http.StatusInternalServerError).WithMessage(err.Error())
}
