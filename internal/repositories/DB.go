package repositories

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
			panic("Cant Open connection to DB")
		}
		db = bun.NewDB(PostGrestDB, pgdialect.New())

		fmt.Println("Connected to Database")

	})
	return db

}
