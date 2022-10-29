package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type repo struct {
	dbClient  *dynamodb.Client
	tableName string
}

func (r *repo) createItem(ctx context.Context, in interface{}) error {
	item, err := attributevalue.MarshalMap(in)
	if err != nil {
		return fmt.Errorf("unable to marshal payload: %w", err)
	}

	_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      item,
	})

	if err != nil {
		return fmt.Errorf("cannot put item: %w", err)
	}

	return nil
}

func (r *repo) getItem(ctx context.Context, input *dynamodb.GetItemInput, output interface{}) error {
	result, err := r.dbClient.GetItem(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to GetItem from DynamoDB: %w", err)
	}

	if len(result.Item) == 0 {
		return nil
	}

	if err := attributevalue.UnmarshalMap(result.Item, output); err != nil {
		return fmt.Errorf("failed to unmarshal data from DynamoDB: %w", err)
	}

	return nil
}

func (r *repo) scan(ctx context.Context, input *dynamodb.ScanInput, output interface{}) error {
	result, err := r.dbClient.Scan(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to Scan from DynamoDB: %w", err)
	}

	if err := attributevalue.UnmarshalListOfMaps(result.Items, output); err != nil {
		return fmt.Errorf("failed to unmarshal data from DynamoDB: %w", err)
	}

	return nil
}

func (r *repo) query(ctx context.Context, input *dynamodb.QueryInput, output interface{}) error {
	result, err := r.dbClient.Query(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to get items from DynamoDB: %w", err)
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, output)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data from DynamoDB: %w", err)
	}

	return nil
}

func (r *repo) updateItem(ctx context.Context, input *dynamodb.UpdateItemInput) error {
	_, err := r.dbClient.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("cannot update item: %w", err)
	}

	return nil
}

func (r *repo) deleteItem(ctx context.Context, input *dynamodb.DeleteItemInput) error {
	_, err := r.dbClient.DeleteItem(ctx, input)

	if err != nil {
		return fmt.Errorf("cannot delete item: %w", err)
	}

	return nil
}
