package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/cliffdoyle/social-network/internal/database"
	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/cliffdoyle/social-network/internal/service"
	_ "github.com/mattn/go-sqlite3"
	_"github.com/golang-migrate/migrate/v4"
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

	// db     *sql.DB
	db *database.DB // Add database connection
	// Add the user service
	services       service.UserService
	sessionService service.SessionService
}

func main() {
	// Declare an instance of the config
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	// Initialize a new structured logger which writes log entries to the standard output stream
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize database connection with your existing database package
	db, err := database.New(database.Config{
		DatabasePath:   "./social_network.db",
		MigrationsPath: "./internal/migrations",
		MaxOpenConns:   25,
		MaxIdleConns:   25,
	})
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	sessionsRepo := repository.NewSessionRepository(db)

	// Initialize Services
	userService := service.NewUserService(userRepo)
	sessionsService := service.NewSessionService(sessionsRepo)

	// Inject dependencies into the application struct
	app := &application{
		config: cfg,
		logger: logger,

		db: db,

		// Inject the service
		services: userService,
		sessionService: sessionsService,
	}

	// Declare new servemux which dispatches requests to
	// our currently single handler method

	// Declare a HTTP server which listens on the port provided in the config struct,
	// uses the servemux created above as the handler and writes any
	// log messages to the structured logger at Error level

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
