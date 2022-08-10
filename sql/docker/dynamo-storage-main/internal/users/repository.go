package users

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

var (
	ErrNotImplemented = fmt.Errorf("metodo no implementado")
)

func itemToUser(input map[string]*dynamodb.AttributeValue) (User, error) {
	var item User
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return User{}, err
	}
	return item, nil
}

type Repository interface {
	Store(ctx context.Context, model *User) error
	GetOne(ctx context.Context, id string) (User, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewRepository(dynamo *dynamodb.DynamoDB, table string) Repository {
	return &repository{
		dynamo: dynamo,
		table:  table,
	}
}

func (receiver *repository) GetOne(ctx context.Context, id string) (User, error) {
	result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return User{}, err
	}

	if result.Item == nil {
		return User{}, nil
	}
	return itemToUser(result.Item)
}

func (receiver *repository) Store(ctx context.Context, model *User) error {
	av, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(receiver.table),
	}

	_, err = receiver.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return ErrNotImplemented
}
