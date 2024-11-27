package main

import (
	"fmt"
	"food-order/internal/repositories"
)

func main() {
	db := repositories.InitDB()
	if db != nil {
		fmt.Println("fine")
	}
}
