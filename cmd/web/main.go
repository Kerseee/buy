package main

import (
	"flag"
	"fmt"
	"log"
)

const version = "1.0.0" // version of this application

type config struct {
	port int    // port of the server listening on
	env  string // development or production
	api  string // url of the backend api
}

type application struct {
	config  config
	version string
}

// serve starts a http server and serves the application.
func (app *application) serve() error {
	router, err := app.routes()
	if err != nil {
		return err
	}
	return router.Run(fmt.Sprintf(":%d", app.config.port))
}

func main() {
	// Populate a config.
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|application}")
	flag.StringVar(&cfg.api, "api", "http://localhost:8081", "Url to api")

	flag.Parse()

	// Create an application.
	app := &application{
		config:  cfg,
		version: version,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
