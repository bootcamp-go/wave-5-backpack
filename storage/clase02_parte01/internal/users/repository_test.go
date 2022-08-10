package users

import (
	"database/sql"
	"fmt"
	"goweb/internal/domain"
	"log"
	"testing"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-playground/assert/v2"
	"golang.org/x/net/context"
)

func TestGetAllUsers(t *testing.T) {

	dataSource := "root:rootpass@tcp(localhost:3306)/storage"
    // Open inicia un pool de conexiones. Sólo abrir una vez
    var err error
    StorageDB, err := sql.Open("mysql", dataSource)
    if err != nil {
        log.Fatal(err)
    }
    if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
    log.Println("database Configured")

    // definimos un users igual al registro de la DB
    users := []domain.User{
		{
            Id: 1,
            Name: "Claudio",
            LastName: "Bieler",
            Email: "taca@gmail.com",
            Age: 39,
            Height: 1.78,
            Active: true,
            CreatedAt: "2005-07-04",
        },
        {
            Id: 2,
            Name: "Lionel",
            LastName: "Messi",
            Email: "lio@gmail.com",
            Age: 35,
            Height: 1.7,
            Active: true,
            CreatedAt: "2015/08/04",
        },
    }
    myRepo := NewRepository(StorageDB)
    // se define un context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    usersResult, err := myRepo.GetAllUsers(ctx)
    fmt.Println(err)
    assert.Equal(t, users[0].Name, usersResult[0].Name)
    assert.Equal(t, users[1].Id, usersResult[1].Id)
}