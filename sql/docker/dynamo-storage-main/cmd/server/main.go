package main

import (
	"log"

	"dynamo-storage-main/cmd/server/handler"
	"dynamo-storage-main/internal/users"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	region := "us-east-1"
	endpoint := "http://localhost:8000"
	cred := credentials.NewStaticCredentials("local", "local", "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	dynamo := dynamodb.New(sess)

	repo := users.NewRepository(dynamo, "users")
	service := users.NewService(repo)
	p := handler.NewUser(service)

	r := gin.Default()
	pr := r.Group("/users")
	pr.GET("/:id", p.GetOne())
	pr.POST("/", p.Store())
	pr.DELETE("/:id", p.Delete())
	r.Run(":8080")
}
