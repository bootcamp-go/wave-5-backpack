# hackaton-go-bases
Repositorio base para hackaton - Go Bases

# uso:
## para crear un ticket

- go run main.go create '{"names":"$names", "email":"$email", "destination":"$destination", "date": "$date", "price": $price}'

## para leer un ticket

- go run main.go read '{"id":$id}'

## para actualizar un ticket

- go run main.go update '{"id":$id, "names":"$names", "email":"$email", "destination":"$destination", "date": "$date", "price": $price}'

## para borrar un ticket

- go run main.go delete '{"id":$id}'