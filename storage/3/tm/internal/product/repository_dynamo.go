package product

import (
	"context"
	"fmt"
	"storage/3/tm/internal/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type RepositoryDynamo interface {
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetOne(ctx context.Context, id string) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id string) error
}

type repositoryDynamo struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewDynamoRepository(dynamo *dynamodb.DynamoDB, table string) RepositoryDynamo {
	return &repositoryDynamo{
		dynamo: dynamo,
		table:  table,
	}
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

func itemToProduct(input map[string]*dynamodb.AttributeValue) (domain.Product, error) {
	var item domain.Product
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return domain.Product{}, err
	}
	return item, nil
}

func (r *repositoryDynamo) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	product.ID = uuid.New().String()

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

	return r.GetOne(ctx, product.ID)
}

func (r *repositoryDynamo) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	result, err := r.dynamo.Scan(&dynamodb.ScanInput{
		TableName: aws.String(r.table),
	})
	if err != nil {
		return []domain.Product{}, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &products)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

func (r *repositoryDynamo) GetOne(ctx context.Context, id string) (domain.Product, error) {
	result, err := r.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return domain.Product{}, err
	}

	if result.Item == nil {
		return domain.Product{}, nil
	}

	return itemToProduct(result.Item)
}

func (r *repositoryDynamo) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	_, err := r.dynamo.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(product.ID),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("name"),
			"#T": aws.String("type"),
			"#C": aws.String("count"),
			"#P": aws.String("price"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": {
				S: aws.String(product.Name),
			},
			":type": {
				S: aws.String(product.Type),
			},
			":count": {
				N: aws.String(fmt.Sprint(product.Count)),
			},
			":price": {
				N: aws.String(fmt.Sprint(product.Price)),
			},
		},
		UpdateExpression: aws.String("SET #N = :name, #T = :type, #C = :count, #P = :price"),
	})
	if err != nil {
		return domain.Product{}, err
	}

	return r.GetOne(ctx, product.ID)
}

func (r *repositoryDynamo) Delete(ctx context.Context, id string) error {
	_, err := r.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
