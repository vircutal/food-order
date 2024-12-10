package utils

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var (
	once sync.Once
	db   *bun.DB
)

func InitDB() *bun.DB {

	once.Do(func() {
		DatabaseSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			"root",
			"root",
			"localhost",
			"5432",
			"food-order-database")
		PostGrestDB, err := sql.Open("postgres", DatabaseSourceName)

		if err != nil {
			//TODO : Print a phase that provides a good meaning
			panic("Cant Open connection to DB")
		}
		db = bun.NewDB(PostGrestDB, pgdialect.New())

		//TODO : Print a phase that provides a good meaning
		fmt.Println("Connected to Database")

	})
	return db
}
