package console

import (
	"net/http"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/config"
	dbmysql "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/db/mysql"
	httphandler "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/delivery/http"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/repository"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		// Init DB connection
		dbConn := dbmysql.InitDBConn()

		// Init repository
		storyRepository := repository.NewStoryRepository(dbConn)
		userRepository := repository.NewUserRepository(dbConn)

		// Init usecae
		storyUsecase := usecase.NewStoryUsecase(storyRepository, userRepository)
		authUsecase := usecase.NewAuthUsecase(userRepository)

		// Run server
		e := echo.New()

		e.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong!")
		})

		storyHandler := httphandler.NewStoryHandler(storyUsecase)
		authHandler := httphandler.NewAuthHandler(authUsecase)

		storyHandler.RegisterRoutes(e)
		authHandler.RegisterRoutes(e)

		e.Start(":" + config.Port())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
