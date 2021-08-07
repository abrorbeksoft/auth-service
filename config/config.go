package config

import (
	"fmt"
	"os"
	"github.com/spf13/cast"
	"github.com/joho/godotenv"
)

type Config struct {
	AuthServiceHost string `json:"auth_service_host"`
	AuthServicePort int `json:"auth_service_port"`
}

func Load() Config {
	if err:=godotenv.Load(); err != nil {
		fmt.Println("No .env file found!")
	}
	config:=Config{}
	config.AuthServiceHost=cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST","localhost"))
	config.AuthServicePort=cast.ToInt(getOrReturnDefault("AUTH_SERVICE_PORT",8001))
	return config
}

func getOrReturnDefault(key string,defaultValue interface{}) interface{} {
	_,exists:=os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}