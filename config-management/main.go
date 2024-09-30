package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	fmt.Println("port", viper.GetInt64("port"))
	fmt.Println("env", viper.GetString("env"))
	fmt.Println("mysql host", viper.GetString("mysql.host"))
	fmt.Println("mysql username", viper.GetString("mysql.username"))
	fmt.Println("mysql password", viper.GetString("mysql.password"))
	fmt.Println("foobar", viper.GetString("foo.bar"))

	fmt.Println(os.Getenv("PORT"))
}
