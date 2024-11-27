package main

import (
	"fmt"
	"food-order/internal/cli"
	"food-order/internal/repositories"
)

func main() {
	db := repositories.InitDB()
	if db != nil {
		fmt.Println("fine")
	}

	if err := cli.RootCmd.Execute(); err != nil {
		return 
	}
}
