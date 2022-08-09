package products

import (
	"context"
	"database/sql"
	"testing"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetOneWithContext(t *testing.T) {
	id := 5
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepo(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = repo.GetOneWithcontext(ctx, id)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}
}
