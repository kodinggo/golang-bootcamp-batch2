package console

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	dbmysql "github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/db/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var (
	direction = "up"
	step      = 1
)

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate is a command to run the migration files",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, err := sql.Open("mysql", dbmysql.NewConnStr())
		if err != nil {
			panic(err)
		}

		dbConn.SetConnMaxLifetime(time.Minute * 3)
		dbConn.SetMaxOpenConns(10)
		dbConn.SetMaxIdleConns(10)

		if err := dbConn.Ping(); err != nil {
			panic(err)
		}

		if len(args) > 0 {
			direction = args[0]
		}

		switch direction {
		case "up", "down":
		default:
			log.Fatal("direction is not valid")
		}

		migrations := &migrate.FileMigrationSource{Dir: "./db/migrations"}
		var n int // "n" berapa migration yg di-applied
		if direction == "down" {
			n, err = migrate.ExecMax(dbConn, "mysql", migrations, migrate.Down, step)
		} else {
			n, err = migrate.ExecMax(dbConn, "mysql", migrations, migrate.Up, step)
		}
		if err != nil {
			log.Fatalf("failed to run migration, error: %s", err.Error())
		}

		log.Printf("successfully applied %d migration(s)", n)
	},
}

func init() {
	migrationCmd.Flags().IntVarP(&step, "step", "s", 1, "number of migration step")
	rootCmd.AddCommand(migrationCmd)
}
