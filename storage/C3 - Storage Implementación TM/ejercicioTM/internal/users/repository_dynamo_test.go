package users

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

//EJERCICIO 3
//Replicar tests de la implementación SQL

//Inicialización de conexión
func InitDynamo() (*dynamodb.DynamoDB, error) {
	region := "us-east-1"
	endpoint := "http://localhost:8000"
	cred := credentials.NewStaticCredentials("local", "local", "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	dynamo := dynamodb.New(sess)
	return dynamo, nil
}

//Testeo de GetOne
func TestGetOneDynamo(t *testing.T) {
	//Obteniendo al user con id "id"
	id := "c65a5be9-ee83-4eb8-b526-0e36efc8d0ec"
	dynamoConn, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamoConn, "users")
	result, err := repo.GetOne(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	//Asserts
	assert.NotNil(t, result)
}

//Testeo del Update
func TestUpdateDynamo(t *testing.T) {
	//Obteniendo al user con id "id"
	id := "c65a5be9-ee83-4eb8-b526-0e36efc8d0ec"
	dynamoConn, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamoConn, "users")
	usuario := User{
		Nombre:   "Luz",
		Apellido: "Lucumi Hernandez",
		Email:    "luz.lucumi@gmail.com",
		Edad:     26,
		Altura:   1.65,
		Activo:   false,
	}
	result, err := repo.Update(id, usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, usuario.Activo)
	if err != nil {
		t.Fatal(err)
	}
	//Asserts
	assert.NotNil(t, result)
}

//Testeo del Delete
func TestDeleteDynamo(t *testing.T) {
	//Eliminando al user con id "id"
	id := "c65a5be9-ee83-4eb8-b526-0e36efc8d0ec"
	dynamoConn, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamoConn, "users")
	delete, err := repo.GetOne(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	result := repo.Delete(context.Background(), id)

	then, err := repo.GetOne(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	//Asserts
	assert.Nil(t, result)
	assert.NotEqual(t, delete, then)
}

//Testeo del Store
func TestStoreDynamo(t *testing.T) {
	//Obteniendo al user con id "id"
	id := "c65a5be9-ee83-4eb8-b526-0e36efc8d0ec"
	dynamoConn, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamoConn, "users")
	usuario := &User{
		Id:       id,
		Nombre:   "Luz",
		Apellido: "Lucumi Hernandez",
		Email:    "luz.lucumi@gmail.com",
		Edad:     26,
		Altura:   1.65,
	}
	errStore := repo.Store(context.TODO(), usuario)
	if errStore != nil {
		t.Fatal(errStore)
	}
	//Asserts
	assert.Nil(t, errStore)
}
