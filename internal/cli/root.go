package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Root",
	Short: "Root command used as an entry point",
	RunE: func(cmd *cobra.Command, args []string) error {
		if(len(args) == 0){
			err := fmt.Errorf("Need an extra argument...")
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(CreateSqlScriptCommand)
}
