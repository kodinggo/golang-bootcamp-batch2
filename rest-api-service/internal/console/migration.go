package console

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate is a command to run the migration files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run the migration files")
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
