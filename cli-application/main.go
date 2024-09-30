package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	},
}

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs application server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("application is running")

		var arg1 string
		if len(args) > 0 {
			arg1 = args[0]
		}
		fmt.Println(arg1)

		fmt.Println("port", port)
	},
}

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate migrates database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("successfully migrate some tables")
	},
}

func main() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "port is ...")

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrationCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
