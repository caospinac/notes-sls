package util

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// EventResponse is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type EventResponse events.APIGatewayV2HTTPResponse
type EventRequest events.APIGatewayV2HTTPRequest

type Handler func(context.Context, EventRequest) (Response, ApiError)
type lambdaHandler func(context.Context, EventRequest) (*EventResponse, error)

func getLambdaHandler(handler Handler) lambdaHandler {
	return func(ctx context.Context, event EventRequest) (*EventResponse, error) {
		eventJSON, _ := json.Marshal(event)
		log.Printf("request_id:%s", event.RequestContext.RequestID)
		log.Printf("event:%s", string(eventJSON))

		res, err := handler(ctx, event)
		if err != nil {

			return err.build(), nil
		}

		return res.build(), nil
	}
}

func Start(handler Handler) {
	lambda.Start(getLambdaHandler(handler))
}
