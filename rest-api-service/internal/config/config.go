package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Setup viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}

func ENV() string {
	return viper.GetString("env")
}

func Port() string {
	return viper.GetString("port")
}

func MySQLDBHost() string {
	return viper.GetString("mysql.dbhost")
}

func MySQLDBUser() string {
	return viper.GetString("mysql.dbuser")
}

func MySQLDBPass() string {
	return viper.GetString("mysql.dbpass")
}

func MySQLDBName() string {
	return viper.GetString("mysql.dbname")
}

func JWTSigningKey() string {
	return viper.GetString("jwt.signing_key")
}

func JWTExp() time.Duration {
	return viper.GetDuration("jwt.exp")
}

func RedisDB() int {
	return viper.GetInt("redis.db")
}

func RedisPass() string {
	return viper.GetString("redis.pass")
}

func RedisHost() string {
	return viper.GetString("redis.host")
}
