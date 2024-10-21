package console

import (
	"log"
	"net"
	"net/http"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/cache"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/config"
	dbmysql "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/db/mysql"
	grpchandler "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/delivery/grpc"
	httphandler "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/delivery/http"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/repository"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/usecase"
	pb "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/pb/rest_api_service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve is a command to run the service server",
	Run: func(cmd *cobra.Command, args []string) {
		// Init DB connection
		dbConn := dbmysql.InitDBConn()
		redisClient := cache.InitRedisClient()

		// Init repository
		storyRepository := repository.NewStoryRepository(dbConn, redisClient)
		userRepository := repository.NewUserRepository(dbConn)

		// Init usecae
		storyUsecase := usecase.NewStoryUsecase(storyRepository, userRepository)
		authUsecase := usecase.NewAuthUsecase(userRepository)

		quitCh := make(chan bool, 1)

		go func() {
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
		}()

		go func() {
			// Run grpc server
			grpcServer := grpc.NewServer()

			storygRPCHandler := grpchandler.NewStorygRPCHandler()

			pb.RegisterStoryServiceServer(grpcServer, storygRPCHandler)

			httpListener, err := net.Listen("tcp", ":6666")
			if err != nil {
				log.Panicf("create http listener: %v", err)
			}

			log.Println("grpc server is running")

			grpcServer.Serve(httpListener)
		}()

		<-quitCh
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
