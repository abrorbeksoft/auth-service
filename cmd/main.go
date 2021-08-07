package main

import (
	"fmt"
	"github.com/abrorbeksoft/auth-service/config"
)

func main()  {
	cfg:=config.Load()
	fmt.Println(cfg.AuthServiceHost)
	fmt.Println(cfg.AuthServicePort)
	fmt.Println("Hello world")
}