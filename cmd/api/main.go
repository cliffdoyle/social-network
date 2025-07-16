package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// Config struct to hold the configuration settings for the application
// For now just the network port that the server will listen on and
// the name of the current operating environment(development, production etc)
// Read these config settings from commandline flags when the application starts
type config struct {
	port int
	env  string
}

// Will hold the dependencies for our http handlers, helpers,
// and middlewares. For now it contains only a copy of the config struct and a logger,but will grow
// to include a lot more as the build progresses
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	//Declare an instance of the config
	var cfg config

	//Read the value of the port and env command-line flags into the config struct.
	//We default to using the port number 4000 and the environment "development" if no
	//corresponding flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	//Initialize a new structured logger which writes log entries to the standard output stream
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Declare an instance of the application struct, containing the config struct and
	//the logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Declare new servemux which dispatches requests to
	//our currently single handler method
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/test", app.errorTest)
	mux.HandleFunc("/api/v1/users/register", app.registerUserHandler)

	//Declare a HTTP server which listens on the port provided in the config struct,
	//uses the servemux created above as the handler and writes any 
	//log messages to the structured logger at Error level

	srv:=&http.Server{
		Addr: fmt.Sprintf(":%d",cfg.port),
		Handler: mux,
		ErrorLog: slog.NewLogLogger(logger.Handler(),slog.LevelError),
	}

	//Start the HTTP server
	logger.Info("starting server","addr",srv.Addr,"env",cfg.env)

	err:=srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
