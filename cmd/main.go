package main

import (
	"eatry/cmd/docs"
	"eatry/pkg/config"
	"eatry/pkg/di"
	"log"
	"os"
)

// @title Eatry - E-commerce API
// @description This is the API documentation for Eatry - E-commerce application.
// @version 1.0
// @host localhost:8000
// @BasePath /

func main() {
	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Eatry - E-commerce"
	docs.SwaggerInfo.Description = "Eatry- E-commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "irin.store"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start(infoLog, errorLog)
	}

}
	