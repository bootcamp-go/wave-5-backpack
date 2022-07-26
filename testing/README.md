## Utilización de Mocks & Stub para testing

Todos los test se encuentran el archivo repository_test.go
Instalación paquete testify
Test de la función restar utilizando testify y la libreria estandar

```go
// import 
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// testify 
go get github.com/stretchr/testify
```

1. Correr test 

```go
// De forma recursiva 
go test ./... -v
```