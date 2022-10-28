package domain

import (
	"github.com/aws/aws-lambda-go/events"
)

// EventResponse is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type EventResponse events.APIGatewayV2HTTPResponse
type EventRequest events.APIGatewayV2HTTPRequest

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
	Offset  *int64      `json:"offset,omitempty"`
	Limit   *int64      `json:"limit,omitempty"`
}

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UpdateBoardRequest struct {
	Title string `json:"title"`
}

type UpdateNoteRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
