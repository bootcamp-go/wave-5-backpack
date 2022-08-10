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

func TestGetOneDynamo(t *testing.T) {
	idSelected := "9db6869d-1cb4-491e-9906-8c49d7985328"
	dynamo, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamo, "users")
	result, err := repo.GetOne(context.Background(), idSelected)
	if err != nil {
		t.Fatal(err)
	}
	// Validation
	assert.NotNil(t, result)
}

func TestDeleteDynamo(t *testing.T) {
	idSelected := "9db6869d-1cb4-491e-9906-8c49d7985328"
	dynamo, err := InitDynamo()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(dynamo, "users")
	beforeDelete, err := repo.GetOne(context.Background(), idSelected)
	if err != nil {
		t.Fatal(err)
	}

	result := repo.Delete(context.Background(), idSelected)

	afterDelete, err := repo.GetOne(context.Background(), idSelected)
	if err != nil {
		t.Fatal(err)
	}

	// Validation
	assert.Nil(t, result)
	assert.NotEqual(t, beforeDelete, afterDelete)
}
