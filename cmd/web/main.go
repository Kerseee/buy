package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0" // version of this application

type config struct {
	port int    // port of the server listening on
	env  string // development or production
	api  string // url of the backend api
	db   struct {
		dsn string // dsn of the database linked by this application
	}
	// stripte stores key and secret needed for linking stripe api
	stripe struct {
		key    string
		secret string
	}
}

type application struct {
	config  config
	version string
	router  *gin.Engine
}

// serve starts a http server and serves the application.
func (app *application) serve() error {
	app.router = gin.Default()
	app.route()
	return app.router.Run(fmt.Sprintf(":%d", app.config.port))
}

func main() {
	// Populate a config.
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|application}")
	flag.StringVar(&cfg.api, "api", "http://localhost:8081", "Url to api")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

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
