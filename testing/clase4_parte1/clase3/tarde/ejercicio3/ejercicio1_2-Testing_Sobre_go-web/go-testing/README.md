#### Gestionar documentación con swaggo

```go
//imports
import (
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)
```

```go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger

// 1.16 o newer
go install github.com/swaggo/swag/cmd/swag@latest
// go mod tidy
```


1. Una vez instalados los paquetes procedemos a generar la documentación

```zsh
swag init -g cmd/server/main.go
```

2. Posteriormente levantamos nuestro servidor y nos vamos a ```http://localhost:8000/docs/index.html``


## Utilización de Stub & Mock para testing

Test de la función restar utilizando testify y la libreria estandar

1. Instalación paquete testify

```go
// import 
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// testify 
go get github.com/stretchr/testify
```

1. Para ejecutar todos los tests

```go
// De forma recursiva 
go test ./... -v 
```

2. Podemos ver 

```go
?       github.com/bootcamp-go/cmd/server     [no test files]
?       github.com/bootcamp-go/cmd/server/handler     [no test files]
?       github.com/bootcamp-go/docs   [no test files]
=== RUN   TestUpdateReceptorYMonto
--- PASS: TestUpdateReceptorYMonto (0.00s)
=== RUN   TestGetAll
--- PASS: TestGetAll (0.00s)
=== RUN   TestGetAllError
--- PASS: TestGetAllError (0.00s)
=== RUN   TestUpdate
--- PASS: TestUpdate (0.00s)
=== RUN   TestDelete
--- PASS: TestDelete (0.00s)
PASS
ok      github.com/bootcamp-go/internal/transactions  0.004s
?       github.com/bootcamp-go/pkg/store      [no test files]
?       github.com/bootcamp-go/pkg/web        [no test files]
```


### Instalación de Go Linter - golangci-lint 

Para instalar este paquete de acuerdo a tu sistema operativo podes chequear la documentación oficial.

https://golangci-lint.run/usage/install/#local-installation

```zsh
// Mi caso linux
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

golangci-lint --version

golangci-lint run
```

1. Ejecutar golangci-lin run y corregir los errores.
2. Medir el coverage del proyecto: 
    - ```go test -cover ./...```  o ```go test -cover ./... -v ``` con información descriptiva

    - Generar reporte: ```go test -cover -coverprofile=coverage.out ./...```
    - Generar reporte con html: ```go tool cover -html=coverage.out``` 

