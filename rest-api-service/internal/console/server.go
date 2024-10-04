package console

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run the server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
