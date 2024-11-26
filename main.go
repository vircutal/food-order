package main

import (
	"fmt"
	"food-order/internal/repositories"
)

func main() {
	repositories.InitDB()
	fmt.Println("fine")
}
