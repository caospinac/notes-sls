package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/caospinac/notes-sls/domain"
)

type Response interface {
	WithBody(interface{}) Response
	WithHeader(string, string) Response
	WithStatus(int) Response
	WithMessage(string) Response

	build() *domain.EventResponse
}

type responseBuilder struct {
	eventResponse domain.EventResponse

	status  int
	message string
	body    interface{}
	headers map[string]string
}

func NewResponse() Response {
	functionName := whoami(1)

	return &responseBuilder{
		status: http.StatusOK,
		headers: map[string]string{
			"Content-Type": "application/json",
			"Func-Reply":   functionName,
		},
	}
}

func (b *responseBuilder) WithBody(in interface{}) Response {
	b.body = in

	return b
}

func (b *responseBuilder) WithHeader(name, value string) Response {
	b.eventResponse.Headers[name] = value

	return b
}

func (b *responseBuilder) WithStatus(code int) Response {
	b.status = code

	return b
}

func (b *responseBuilder) WithMessage(message string) Response {
	b.message = message

	return b
}

func (b *responseBuilder) build() *domain.EventResponse {
	if b.message == "" {
		b.message = http.StatusText(b.status)
	}

	responseBody := &domain.Response{
		Status:  b.status,
		Message: b.message,
	}

	if b.body != nil {
		responseBody.Body = b.body
	}

	responseBytes, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}

	var responseBuffer bytes.Buffer
	json.HTMLEscape(&responseBuffer, responseBytes)

	return &domain.EventResponse{
		StatusCode:      b.status,
		IsBase64Encoded: false,
		Headers:         b.headers,
		Body:            responseBuffer.String(),
	}
}

func whoami(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown"
	}

	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}

	return me.Name()
}
