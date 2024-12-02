package main

import (
	"fmt"
	"food-order/internal/cli"
	"food-order/internal/repositories"
)

// func executeSQLScript(db *bun.DB, scriptPath string) error {
// 	// Read the SQL script file
// 	sqlScript, err := os.ReadFile(scriptPath)
// 	if err != nil {
// 		return fmt.Errorf("error reading SQL script: %w", err)
// 	}

// 	// Execute the SQL script
// 	_, err = db.Exec(string(sqlScript))
// 	if err != nil {
// 		return fmt.Errorf("error executing SQL script: %w", err)
// 	}

// 	return nil
// }

func main() {
	db := repositories.InitDB()
	if db != nil {
		//TODO : Print a phase that provides a good meaning
		fmt.Println("fine")
	}
	defer db.Close()

	// err := executeSQLScript(db, "./internal/migrations/20241127224026_test.up.sql")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
