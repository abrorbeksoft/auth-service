package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	AuthServiceHost string `json:"auth_service_host"`
	AuthServicePort int `json:"auth_service_port"`

	DBHost string `json:"database_host"`
	DBPort int `json:"database_port"`
	DBDatbase string `json:"db_datbase"`
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
}

func Load() Config {
	if err:=godotenv.Load(); err != nil {
		fmt.Println("No .env file found!")
	}
	config:=Config{}
	config.AuthServiceHost=cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST","localhost"))
	config.AuthServicePort=cast.ToInt(getOrReturnDefault("AUTH_SERVICE_PORT",8001))

	config.DBHost=cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	config.DBPort=cast.ToInt(getOrReturnDefault("POSTGRES_PORT",5672))
	config.DBDatbase=cast.ToString(getOrReturnDefault("POSTGRES_DB","auth-service"))
	config.DBUsername=cast.ToString(getOrReturnDefault("POSTGRES_USER","user"))
	config.DBPassword=cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","password"))

	return config
}

func getOrReturnDefault(key string,defaultValue interface{}) interface{} {
	_,exists:=os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}