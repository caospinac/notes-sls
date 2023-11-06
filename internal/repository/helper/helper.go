package helper

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/caospinac/notes-sls/pkg/util"
)

func SingleResult(result *mongo.SingleResult, v interface{}) util.ApiError {
	if err := result.Err(); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return util.NewApiError(http.StatusNotFound)
		}

		return util.ToApiError(err)
	}

	if v != nil {
		if err := result.Decode(v); err != nil {
			return util.ToApiError(err)
		}
	}

	return nil
}

func FilterID(ID interface{}) bson.M {
	return bson.M{"_id": ID}
}

func WithSession(ctx context.Context, client *mongo.Client, fn func() error) util.ApiError {
	session, err := client.StartSession()
	if err != nil {
		return util.ToApiError(err)
	}

	defer session.EndSession(ctx)

	if err := session.StartTransaction(); err != nil {
		return util.ToApiError(err)
	}

	if err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := fn(); err != nil {
			return err
		}

		return sc.CommitTransaction(ctx)
	}); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return util.NewApiError(http.StatusNotFound)
		}

		return util.ToApiError(err)
	}

	return nil
}
