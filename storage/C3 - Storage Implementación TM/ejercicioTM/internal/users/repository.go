package users

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	Id       string    `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

var (
	ErrNotImplemented = fmt.Errorf("método no implementado")
)

//Funciona con GetOne
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
	Update(id string, nombre string, apellido string, email string, edad int, altura float64, activo bool) (User, error)
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

func (receiver *repository) Update(id string, nombre string, apellido string, email string, edad int, altura float64, activo bool) (User, error) {

	type UserUpdate struct {
		Nombre   string  `json:"nombre"`
		Apellido string  `json:"apellido"`
		Email    string  `json:"email"`
		Edad     int     `json:"edad"`
		Altura   float64 `json:"altura"`
		Activo   bool    `json:"activo"`
	}

	updateData, err := dynamodbattribute.MarshalMap(UserUpdate{
		Nombre:   nombre,
		Apellido: apellido,
		Email:    email,
		Edad:     edad,
		Altura:   altura,
		Activo:   activo,
	})

	if err != nil {
		return User{}, err
	}

	updateString := "SET nombre = :no, apellido = :ap, email = :em, edad = :ed, altura = :al, activo = :ac"

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
		return User{}, err
	}
	return itemToUser(result.Attributes)
}
