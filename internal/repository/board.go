package repository

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/caospinac/notes-sls/internal/defines"
	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/repository/helper"
	"github.com/caospinac/notes-sls/pkg/util"
)

type BoardsRepo interface {
	Insert(context.Context, *domain.Board) (string, util.ApiError)
	FindOne(context.Context, string) (*domain.Board, util.ApiError)
	Update(context.Context, string, *domain.Board) util.ApiError
	Delete(context.Context, string) util.ApiError
}

type boardsRepo struct {
	collection      *mongo.Collection
	notesCollection *mongo.Collection
}

func NewBoardsRepo(dbClient *mongo.Database) BoardsRepo {
	collection := dbClient.Collection(defines.CollectionBoards)
	notesCollection := dbClient.Collection(defines.CollectionNotes)
	return &boardsRepo{
		collection,
		notesCollection,
	}
}

func (repo boardsRepo) Insert(ctx context.Context, document *domain.Board) (string, util.ApiError) {
	document.ID = primitive.NewObjectID()
	result, err := repo.collection.InsertOne(ctx, document)
	if err != nil {
		return "", util.ToApiError(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo boardsRepo) FindOne(ctx context.Context, ID string) (*domain.Board, util.ApiError) {
	objectID, errObjectID := primitive.ObjectIDFromHex(ID)
	if errObjectID != nil {
		return nil, util.NewApiError(http.StatusBadRequest)
	}

	result := repo.collection.FindOne(ctx, helper.FilterID(objectID))
	document := new(domain.Board)
	if err := helper.SingleResult(result, document); err != nil {
		return nil, err
	}

	return document, nil
}

func (repo boardsRepo) Update(ctx context.Context, ID string, update *domain.Board) util.ApiError {
	objectID, errObjectID := primitive.ObjectIDFromHex(ID)
	if errObjectID != nil {
		return util.NewApiError(http.StatusBadRequest)
	}

	result := repo.collection.FindOneAndUpdate(ctx, helper.FilterID(objectID), helper.Set(update))

	return helper.SingleResult(result, nil)
}

func (repo boardsRepo) Delete(ctx context.Context, ID string) util.ApiError {
	objectID, errObjectID := primitive.ObjectIDFromHex(ID)
	if errObjectID != nil {
		return util.NewApiError(http.StatusBadRequest)
	}

	err := helper.WithSession(ctx, repo.collection.Database().Client(), func() error {
		result := repo.collection.FindOneAndDelete(ctx, helper.FilterID(objectID))
		deletedBoard := new(domain.Board)
		if err := helper.SingleResult(result, deletedBoard); err != nil {
			return err.Error()
		}

		notesFilter := helper.FilterID(bson.M{"$in": deletedBoard.NoteIDs})
		_, err := repo.notesCollection.DeleteMany(ctx, notesFilter)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
