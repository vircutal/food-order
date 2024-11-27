package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func InitDB() *bun.DB {
	var err error

	DatabaseSourceName := ""
	PostGrestDB, err := sql.Open("postgres", DatabaseSourceName)

	if err != nil {
		panic("Cant Open connection to DB")
	}
	db := bun.NewDB(PostGrestDB, pgdialect.New())
	defer db.Close()

	fmt.Println("Connected to Database")

	return db

}
