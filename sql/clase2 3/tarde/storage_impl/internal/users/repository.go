package users

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bootcamp-go/storage/internal/domains"
	"github.com/google/uuid"
)

type RepositoryDynamo interface {
	Store(context.Context, *domains.User) error
	GetOne(context.Context, string) (*domains.User, error)
	Delete(context.Context, string) error
	Update(context.Context, *domains.User) error
}

type repositoryDynamo struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewDynamoRepository(db *dynamodb.DynamoDB, table string) RepositoryDynamo {
	return &repositoryDynamo{
		dynamo: db,
		table:  table,
	}
}

/* Ejercicio 1 - Implementar Store() y GetOne()
Basándose en el material visto en la clase,
    • Crear una implementación de la interfaz Repository utilizando el sdk de aws de go para desarrollar los métodos de Store y GetOne.
Esto puede ser testeado utilizando dynamodb admin (utilizando docker-compose) o mediante tests. */
func (r *repositoryDynamo) Store(ctx context.Context, u *domains.User) error {
	u.Id = uuid.New().String()
	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}
	_, err = r.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

/* func (r *repositoryDynamo) GetOne(ctx context.Context, id string) (*domains.User, error) {
	result, err := r.dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return domains.ItemToUser(result.Item)
} */

func (r *repositoryDynamo) GetOne(ctx context.Context, id string) (*domains.User, error) {
	result, err := r.dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return domains.ItemToUser(result.Item)
}

/* Ejercicio 2 - Implementar Update() y Delete()
Utilizando la documentación del sdk de go de aws implementar los métodos Update y Delete en DynamoDB. */
func (r *repositoryDynamo) Update(ctx context.Context, u *domains.User) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#F":   aws.String("firstname"),
			"#L":   aws.String("lastname"),
			"#U":   aws.String("username"),
			"#P":   aws.String("password"),
			"#E":   aws.String("email"),
			"#IP":  aws.String("ip"),
			"#MC":  aws.String("macAddress"),
			"#W":   aws.String("website"),
			"#IMG": aws.String("image"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f":   {S: aws.String(u.Firstname)},
			":l":   {S: aws.String(u.Lastname)},
			":u":   {S: aws.String(u.Username)},
			":p":   {S: aws.String(u.Password)},
			":e":   {S: aws.String(u.Email)},
			":ip":  {S: aws.String(u.IP)},
			":mc":  {S: aws.String(u.MacAddress)},
			":w":   {S: aws.String(u.Website)},
			":img": {S: aws.String(u.Image)},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(u.Id)},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		TableName:        aws.String(r.table),
		UpdateExpression: aws.String("SET #F = :f, #L = :l, #U = :u, #P = :p, #E = :e, #IP = :ip, #MC = :mc, #W = :w, #IMG = :img"),
	}

	_, err := r.dynamo.UpdateItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryDynamo) Delete(ctx context.Context, id string) error {
	_, err := r.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
