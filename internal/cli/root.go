package cli

import (
	"food-order/internal/controller"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Root",
	Short: "Root command used as an entry point for starting application",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := controller.GetRootController()

		if err := r.Listen(":3000"); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(CreateSqlScriptCommand)
}
