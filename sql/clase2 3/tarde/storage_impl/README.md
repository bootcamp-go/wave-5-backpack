### DynamoDB

Guía del Desarrollador:

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html
https://aws.github.io/aws-sdk-go-v2/docs/code-examples/dynamodb/

https://github.com/aws/aws-sdk-go

Local-DynamoDB: http://localhost:8000
Admin-DynamoDB: http://localhost:8001

Documentación:
https://pkg.go.dev/github.com/aws/aws-sdk-go@v1.44.70/service/dynamodb


Mock para dynamo: 
https://pkg.go.dev/github.com/gusaul/go-dynamock#section-readme


[POST][http://localhost:8080/api/v1/users]
```json
{
	"firstname":"gregory",
	"lastname":"house",
	"username":"dr. house",
	"password":"asd%Sara#sadf",
	"email":"gregory@digitalhouse.com",
	"ip":"134.532.234.12",
	"macAddress":"AD:DF:FD:fDF:DF:FD",
	"website":"drhouse.com",
	"image":"imagerandom"
}
```

[GET][http://localhost:8080/api/v1/users/37f07ac3-6088-4738-bc93-72288e0112c5] 