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

//Para manipulación por item tipo User
func itemToUser(input map[string]*dynamodb.AttributeValue) (User, error) {
	var item User
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return User{}, err
	}
	return item, nil
}

//Repository con sus métodos
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

//EJERCICIO 1
//Implementar Store() y GetOne()

//Implementación GetOne()
func (receiver *repository) GetOne(ctx context.Context, id string) (User, error) {
	result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		//Pasamos nombre de la tabla y el id
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
	//Si no ocurren errores, devuelve el usuario
	return itemToUser(result.Item)
}

//Implementación Store()
func (receiver *repository) Store(ctx context.Context, model *User) error {
	av, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	//Paso item que será del tipo de la estructura de arriba y el nombre de la tabla
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(receiver.table),
	}

	//Se agrega el elemento
	_, err = receiver.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

//EJERCICIO 2
//Implementar Update() y Delete()

//Implementación del Delete
func (receiver *repository) Delete(ctx context.Context, id string) error {
	//Elimino el elemento por id
	result, err := receiver.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	//Manejo de errores
	if err != nil {
		return err
	}

	if result.Attributes == nil {
		return nil
	}

	return fmt.Errorf("error al eliminar elemento")
}

//Implementación del Update
func (receiver *repository) Update(id string, nombre string, apellido string, email string, edad int, altura float64, activo bool) (User, error) {

	//Struct para asignación en el SET del Update
	type UserUpdate struct {
		Nombre   string  `json:":no"`
		Apellido string  `json:":ap"`
		Email    string  `json:":em"`
		Edad     int     `json:":ed"`
		Altura   float64 `json:":al"`
		Activo   bool    `json:":ac"`
	}

	//Map con información actual
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
		TableName: aws.String(receiver.table),
		//Asignación de valores ingresados al updateString
		ExpressionAttributeValues: updateData,
		UpdateExpression:          aws.String(updateString),
		//Retorna atributos actualizados, tal como aparecen después de la operación UpdateItem
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	//Con UpdateItem se editan los atributos del elemento existente
	result, err := receiver.dynamo.UpdateItem(input)
	if err != nil {
		return User{}, err
	}
	return itemToUser(result.Attributes)
}
