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
	boardsTableName = os.Getenv("DYNAMODB_TABLE_BOARDS")
)

type BoardRepo interface {
	Create(context.Context, *domain.Board) error
	Get(context.Context, string) (*domain.Board, error)
	GetAll(context.Context) ([]domain.Board, error)
	Update(context.Context, string, domain.UpdateBoardRequest) error
	Delete(context.Context, string) error
}

func NewBoardRepo(dbClient *dynamodb.Client) BoardRepo {
	return &boardRepo{
		repo{
			dbClient,
			boardsTableName,
		},
	}
}

type boardRepo struct {
	repo
}

func (r *boardRepo) Create(ctx context.Context, board *domain.Board) error {
	board.ID = helper.NewUniqueID()

	return r.createItem(ctx, board)
}

func (r *boardRepo) Get(ctx context.Context, id string) (*domain.Board, error) {
	board := new(domain.Board)
	if err := r.getItem(ctx, domain.Board{ID: id}, &board); err != nil {
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
	filter := expression.AttributeExists(expression.Name("id"))
	update := expression.Set(expression.Name("title"), expression.Value(&types.AttributeValueMemberS{Value: board.Title}))

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
			"id": &types.AttributeValueMemberS{Value: id},
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

func (r *boardRepo) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}

	return r.deleteItem(ctx, input)
}
