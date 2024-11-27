package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var MigrateCMD = &cobra.Command{
	Use: "migrate",
}

var name string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a greeting message",
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Hello, world!")
		}
	},
}

func Migrate() {
	var err error
	timestamp := time.Now().Format("20060102150405")
	fileName := "test"

	content := []byte("--SQL migration content here")

	finalFileName := timestamp + "_" + fileName
	//create up file
	err = os.WriteFile("internal/migrations/"+finalFileName+".up.sql", content, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//create down file
	err = os.WriteFile("internal/migrations/"+finalFileName+".down.sql", content, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migration file created : ", finalFileName)
}
