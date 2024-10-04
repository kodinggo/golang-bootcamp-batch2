package console

import (
	"log"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "restapi-service",
	Short: "restapi-service is a service for REST API",
}

func init() {
	config.InitConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
