package util

import (
	"net/http"

	"github.com/caospinac/notes-sls/domain"
)

type ApiError interface {
	WithStatus(int) ApiError
	WithMessage(string) ApiError

	build() *domain.EventResponse
}

type apiErrorBuilder struct {
	status  int
	message string
}

func NewApiError() ApiError {
	return &apiErrorBuilder{
		status: http.StatusInternalServerError,
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

func (b *apiErrorBuilder) build() *domain.EventResponse {
	return NewResponse().
		WithStatus(b.status).
		WithMessage(b.message).
		build()
}

func ToApiError(err error) ApiError {
	return NewApiError().WithMessage(err.Error())
}
