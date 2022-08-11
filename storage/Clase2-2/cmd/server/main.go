package main

import (
	"storage/pkg/db"
)

func main() {
	db.StorageDB.Ping()

	// repo := products.NewRepo()
	// prod, err := repo.GetByName("test")
	// prod := domain.Product{Name: "test4", Type: "", Count: 0, Price: 0}
	// result, err := repo.Store(prod)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(result)
}
