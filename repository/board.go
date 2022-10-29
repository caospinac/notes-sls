package repository

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/caospinac/notes-sls/db"
	"github.com/caospinac/notes-sls/domain"
)

var (
	boardTableName = os.Getenv("BOARDS_TABLE")
)

type BoardRepo interface {
	Create(context.Context, domain.Board) error
	Get(context.Context, string) (*domain.Board, error)
	GetAll(context.Context) ([]domain.Board, error)
	Update(context.Context, string, domain.UpdateBoardRequest) error
	Delete(context.Context, string) error
}

func NewBoardRepo() BoardRepo {
	return &boardRepo{
		repo{
			dbClient:  db.NewDynamoDBClient(),
			tableName: boardTableName,
		},
	}
}

type boardRepo struct {
	repo
}

func (r *boardRepo) Create(ctx context.Context, board domain.Board) error {
	return r.createItem(ctx, &board)
}

func (r *boardRepo) Get(ctx context.Context, id string) (*domain.Board, error) {
	input := &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}
	board := new(domain.Board)
	if err := r.getItem(ctx, input, &board); err != nil {
		return nil, err
	}

	if board.ID == "" {
		return nil, nil
	}

	return board, nil
}

func (r *boardRepo) GetAll(ctx context.Context) ([]domain.Board, error) {
	var limit int32 = 20 // TODO: pagination
	input := &dynamodb.ScanInput{
		TableName: &r.tableName,
		Limit:     &limit,
	}

	boards := make([]domain.Board, 0)
	if err := r.scan(ctx, input, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *boardRepo) Update(ctx context.Context, id string, board domain.UpdateBoardRequest) error {
	input := &dynamodb.UpdateItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		AttributeUpdates: map[string]types.AttributeValueUpdate{
			"title": {
				Action: types.AttributeActionPut,
				Value: &types.AttributeValueMemberS{
					Value: board.Title,
				},
			},
		},
	}

	if err := r.updateItem(ctx, input); err != nil {
		return err
	}

	return nil
}

func (r *boardRepo) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}

	return r.deleteItem(ctx, input)
}
