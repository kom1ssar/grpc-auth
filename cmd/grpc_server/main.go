package main

import (
	"context"
	"flag"
	"github.com/kom1ssar/grpc-auth/internal/app"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "local.env", "path to config file")
}

func main() {
	a, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("erro to init app %+v", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("erro to run app %+v", err.Error())

	}
}
