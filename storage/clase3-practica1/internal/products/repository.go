package products

import (
	"clase3-practica1/internal/domain"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func itemToProduct(item map[string]*dynamodb.AttributeValue) (*domain.Product, error) {
	var product domain.Product

	err := dynamodbattribute.UnmarshalMap(item, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

type Repository interface {
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	GetOne(ctx context.Context, id int) (*domain.Product, error)
	//Update(ctx context.Context, product domain.Product) (domain.Product, error)
	//Delete(ctx context.Context, id int) error
}

type repository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewRepository(dynamo *dynamodb.DynamoDB, table string) Repository {
	return &repository{dynamo, table}
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	av, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return domain.Product{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}

	_, err = r.dynamo.PutItemWithContext(ctx, input)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) GetOne(ctx context.Context, id int) (*domain.Product, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%d", id)),
			},
		},
	}

	result, err := r.dynamo.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("product not found")
	}

	return itemToProduct(result.Item)
}
