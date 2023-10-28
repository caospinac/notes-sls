package repository

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/caospinac/notes-sls/internal/domain"
	"github.com/caospinac/notes-sls/internal/repository/helper"
)

var (
	notesTableName = os.Getenv("DYNAMODB_TABLE_NOTES")
)

type NoteRepo interface {
	Create(context.Context, string, *domain.Note) error
	GetAll(context.Context, string) ([]domain.Note, error)
	Update(context.Context, string, string, domain.UpdateNoteRequest) error
	Delete(context.Context, string, string) error
}

type noteRepo struct {
	repo
}

func NewNoteRepo(dbClient *dynamodb.Client) NoteRepo {
	return &noteRepo{
		repo{
			dbClient,
			notesTableName,
		},
	}
}

func (r *noteRepo) Create(ctx context.Context, boardID string, note *domain.Note) error {
	note.BoardID = boardID
	note.NoteID = helper.NewUniqueID()

	return r.createItem(ctx, &note)
}

func (r *noteRepo) GetAll(ctx context.Context, boardID string) ([]domain.Note, error) {
	var limit int32 = 20 // TODO: pagination && filtering

	input := &dynamodb.QueryInput{
		TableName: &r.tableName,
		Limit:     &limit,
		KeyConditions: map[string]types.Condition{
			"board_id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{
						Value: boardID,
					},
				},
			},
		},
	}

	notes := make([]domain.Note, 0)
	if err := r.query(ctx, input, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *noteRepo) Update(ctx context.Context, boardID, noteID string, note domain.UpdateNoteRequest) error {
	filter := expression.And(
		expression.AttributeExists(expression.Name("board_id")), expression.AttributeExists(expression.Name("note_id")),
	)
	update := expression.
		Set(expression.Name("title"), expression.Value(&types.AttributeValueMemberS{Value: note.Title})).
		Set(expression.Name("description"), expression.Value(&types.AttributeValueMemberS{Value: note.Description}))

	expr, err := expression.NewBuilder().
		WithCondition(filter).
		WithUpdate(update).
		Build()

	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"board_id": &types.AttributeValueMemberS{Value: boardID},
			"note_id":  &types.AttributeValueMemberS{Value: noteID},
		},
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	if err := r.updateItem(ctx, input); err != nil {
		return err
	}

	return nil
}

func (r *noteRepo) Delete(ctx context.Context, boardID, noteID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"board_id": &types.AttributeValueMemberS{Value: boardID},
			"note_id":  &types.AttributeValueMemberS{Value: noteID},
		},
	}

	return r.deleteItem(ctx, input)
}
