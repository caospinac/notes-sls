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

type NotesRepo interface {
	Insert(ctx context.Context, boardID string, document *domain.Note) (string, util.ApiError)
	FindOne(context.Context, string) (*domain.Note, util.ApiError)
	FindByBoardID(ctx context.Context, boardID string) ([]domain.Note, util.ApiError)
	Update(context.Context, string, *domain.Note) util.ApiError
	Delete(context.Context, string) util.ApiError
}

type notesRepo struct {
	collection       *mongo.Collection
	boardsCollection *mongo.Collection
}

func NewNotesRepo(dbClient *mongo.Database) NotesRepo {
	collection := dbClient.Collection(defines.CollectionNotes)
	boardsCollection := dbClient.Collection(defines.CollectionBoards)
	return &notesRepo{
		collection,
		boardsCollection,
	}
}

func (repo notesRepo) Insert(ctx context.Context, boardID string, document *domain.Note) (string, util.ApiError) {
	document.ID = primitive.NewObjectID()
	boardObjectID, errObjectID := primitive.ObjectIDFromHex(boardID)
	if errObjectID != nil {
		return "", util.NewApiError(http.StatusBadRequest)
	}

	var result *mongo.InsertOneResult
	document.BoardID = boardObjectID

	err := helper.WithSession(ctx, repo.collection.Database().Client(), func() error {
		insertResult, err := repo.collection.InsertOne(ctx, document)
		if err != nil {
			return err
		}

		update := bson.M{
			"$addToSet": bson.M{"notes": insertResult.InsertedID},
		}
		updatedResult := repo.boardsCollection.FindOneAndUpdate(ctx, helper.FilterID(boardObjectID), update)
		if updatedResult.Err() != nil {
			return err
		}

		result = insertResult

		return nil
	})

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo notesRepo) FindOne(ctx context.Context, ID string) (*domain.Note, util.ApiError) {
	objectID, errObjectID := primitive.ObjectIDFromHex(ID)
	if errObjectID != nil {
		return nil, util.NewApiError(http.StatusBadRequest)
	}

	result := repo.collection.FindOne(ctx, helper.FilterID(objectID))
	document := new(domain.Note)
	if err := helper.SingleResult(result, document); err != nil {
		return nil, err
	}

	return document, nil
}

func (repo notesRepo) FindByBoardID(ctx context.Context, boardID string) ([]domain.Note, util.ApiError) {
	boardObjectID, errObjectID := primitive.ObjectIDFromHex(boardID)
	if errObjectID != nil {
		return nil, util.NewApiError(http.StatusBadRequest)
	}

	filter := bson.M{"board_id": boardObjectID}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, util.ToApiError(err)
	}

	documents := make([]domain.Note, 0)
	cursor.All(ctx, &documents)

	return documents, nil
}

func (repo notesRepo) Update(ctx context.Context, ID string, update *domain.Note) util.ApiError {
	objectID, errObjectID := primitive.ObjectIDFromHex(ID)
	if errObjectID != nil {
		return util.NewApiError(http.StatusBadRequest)
	}

	result := repo.collection.FindOneAndUpdate(ctx, helper.FilterID(objectID), update)

	return helper.SingleResult(result, nil)
}

func (repo notesRepo) Delete(ctx context.Context, noteID string) util.ApiError {
	noteObjectID, errObjectID := primitive.ObjectIDFromHex(noteID)
	if errObjectID != nil {
		return util.NewApiError(http.StatusBadRequest)
	}

	err := helper.WithSession(ctx, repo.collection.Database().Client(), func() error {
		result := repo.collection.FindOneAndDelete(ctx, helper.FilterID(noteObjectID))
		deletedNote := new(domain.Note)
		if err := helper.SingleResult(result, deletedNote); err != nil {
			return err.Error()
		}

		update := bson.M{
			"$pull": bson.M{"notes": deletedNote.ID},
		}
		updatedResult := repo.boardsCollection.FindOneAndUpdate(ctx, helper.FilterID(deletedNote.BoardID), update)
		if err := updatedResult.Err(); err != nil {
			return err
		}

		return nil
	})

	return err
}
