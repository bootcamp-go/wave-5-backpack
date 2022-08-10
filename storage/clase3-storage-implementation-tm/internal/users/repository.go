package users

import (
	"context"
	"fmt"

	"clase3-storage-implementation-tm/internal/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ErrNotImplemented | Variable ERRORS
var (
	ErrNotImplemented = fmt.Errorf("metodo no implementado")
)

// Repository ...
type Repository interface {
	Store(ctx context.Context, model *domain.User) error
	GetOne(ctx context.Context, id string) (domain.User, error)
	Update(id, firstname, lastname, username, email string) (domain.User, error)
	Delete(ctx context.Context, id string) error
}

// repository ...
type repository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

// NewRepository ...
func NewRepository(dynamo *dynamodb.DynamoDB, table string) Repository {
	return &repository{
		dynamo: dynamo,
		table:  table,
	}
}

func itemToUser(input map[string]*dynamodb.AttributeValue) (domain.User, error) {
	var item domain.User
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return domain.User{}, err
	}
	return item, nil
}

func (receiver *repository) GetOne(ctx context.Context, id string) (domain.User, error) {
	result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return domain.User{}, err
	}

	if result.Item == nil {
		return domain.User{}, nil
	}
	return itemToUser(result.Item)
}

func (receiver *repository) Store(ctx context.Context, model *domain.User) error {
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

// Update ...
func (receiver *repository) Update(id, firstname, lastname, username, email string) (domain.User, error) {

	type UserUpdate struct {
		Firstname string `json:":fn"`
		Lastname  string `json:":ln"`
		Username  string `json:":un"`
		Email     string `json:":em"`
	}

	updateData, err := dynamodbattribute.MarshalMap(UserUpdate{
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Email:     email,
	})

	if err != nil {
		return domain.User{}, err
	}

	updateString := "SET firstname = :fn, lastname = :ln, username = :un, email = :em"

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName:                 aws.String(receiver.table),
		ExpressionAttributeValues: updateData,
		UpdateExpression:          aws.String(updateString),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	result, err := receiver.dynamo.UpdateItem(input)
	if err != nil {
		return domain.User{}, err
	}
	return itemToUser(result.Attributes)
}

func (receiver *repository) Delete(ctx context.Context, id string) error {
	result, err := receiver.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return err
	}

	if result.Attributes == nil {
		return nil
	}

	return fmt.Errorf("error al eliminar elemento")
}
