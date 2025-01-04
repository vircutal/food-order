package cli

import (
	"fmt"
	"food-order/internal/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	RunE: func(cmd *cobra.Command, args []string) error {
		migrationDir := "./internal/migrations/"
		var upFiles []string

		files, err := os.ReadDir(migrationDir)
		if err != nil {
			fmt.Println("Cannot read migrations directory...")
			return err
		}

		for _, val := range files {
			if !val.IsDir() && strings.HasSuffix(val.Name(), ".up.sql") {
				upFiles = append(upFiles, val.Name())
			}
		}

		db := utils.InitDB()
		fmt.Println("Migrate these file")
		for _, val := range upFiles {
			script, err := os.ReadFile(migrationDir + val)
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = db.Exec(string(script))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(" - " + val)
		}
		return nil
	},
}
