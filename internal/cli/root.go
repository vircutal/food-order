package cli

import (

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Root",
	Short: "Root command used as an entry point",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(CreateSqlScriptCommand)
}
