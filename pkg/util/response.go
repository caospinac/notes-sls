package util

type Response interface {
	WithBody(interface{}) Response
	WithHeader(string, string) Response

	build() *EventResponse
}

type responseBuilder struct {
	status  int
	body    interface{}
	headers map[string]string
}

func NewResponse(status int) Response {
	return &responseBuilder{
		status: status,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (b *responseBuilder) WithBody(in interface{}) Response {
	b.body = in

	return b
}

func (b *responseBuilder) WithHeader(name, value string) Response {
	b.headers[name] = value

	return b
}

func (b *responseBuilder) build() *EventResponse {

	return &EventResponse{
		StatusCode:      b.status,
		IsBase64Encoded: false,
		Headers:         b.headers,
		Body:            getBodyString(b.body),
	}
}
