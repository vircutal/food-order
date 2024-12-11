package main

import (
	"fmt"
	"food-order/internal/cli"
	"food-order/internal/utils"
)

func main() {
	//InitDB here so it can be shut down properly
	db := utils.InitDB()
	if db != nil {
		//TODO : Print a phase that provides a good meaning
		fmt.Println("fine")
	}
	defer db.Close()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
