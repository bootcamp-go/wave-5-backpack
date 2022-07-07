## Instalar gin en el projecto

Previamente necesitamos crear un repositorio
1. Crear un repositorio
2. Clonar el repositorio

```go
go mod init github.com/usernamegithub/go-web
// go 1.14+
go get -u github.com/gin-gonic/gin
go mod tidy
```

3. Una vez implementado gin - y el endpoint
4. Ejecutamos nuestro main

```go
go run main.go
```


5. Extra como obtener la info de la solicitud http
```curl
curl -X GET http://localhost:8000/nombre -I
```

