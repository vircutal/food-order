package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var CreateSqlScriptCommand = &cobra.Command{
	Use:   "new-sql",
	Short: "Creating Sql Script",
	RunE: func(cmd *cobra.Command, args []string) error{
		if(len(args) == 0){
			err := fmt.Errorf("Need an extra argument...")
			return err
		}
		fileName := args[0]
		var err error
		timestamp := time.Now().Format("20060102150405")

		content := []byte("--SQL migration content here")

		finalFileName := timestamp + "_" + fileName
		//create up file
		err = os.WriteFile("internal/migrations/"+finalFileName+".up.sql", content, 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}

		//create down file
		err = os.WriteFile("internal/migrations/"+finalFileName+".down.sql", content, 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Migration file created : ", finalFileName)

		return nil
	},
}
