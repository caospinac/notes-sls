package util

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/caospinac/notes-sls/domain"
)

type Handler func(context.Context, domain.EventRequest) (Response, ApiError)
type lambdaHandler func(context.Context, domain.EventRequest) (*domain.EventResponse, error)

func getLambdaHandler(handler Handler) lambdaHandler {
	return func(ctx context.Context, event domain.EventRequest) (*domain.EventResponse, error) {
		eventJSON, _ := json.Marshal(event)
		log.Printf("event:%s", string(eventJSON))

		res, err := handler(ctx, event)
		if err != nil {
			log.Printf("request_id:%s, error:%s", event.RequestContext.RequestID, err.getMessage())

			return err.build(), nil
		}

		return res.build(), nil
	}
}

func Start(handler Handler) {
	lambda.Start(getLambdaHandler(handler))
}
