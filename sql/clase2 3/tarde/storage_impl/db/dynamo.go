package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitDynamo() *dynamodb.DynamoDB {
	region := "us-west-2"
	endpoint := "http://localhost:8000"
	cred := credentials.NewStaticCredentials("local", "local", "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))
	if err != nil {
		panic(err)
	}

	dynamo := dynamodb.New(sess)
	return dynamo
}
