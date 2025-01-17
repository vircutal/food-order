package main

import (
	"fmt"
	"food-order/internal/cli"
	"food-order/internal/utils"
)

// b, _ := db.Begin()
// btx, _ := db.BeginTx()
func main() {
	//InitDB here so it can be shut down properly
	db := utils.InitDB()
	if db == nil {
		//TODO : Print a phase that provides a good meaning
		fmt.Println("Error with initialize db")
	}
	defer db.Close()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
