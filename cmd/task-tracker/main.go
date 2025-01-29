package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"taskTracker/constants"
	"taskTracker/core/iniservice"
	"taskTracker/core/relationaldatabase"
	"taskTracker/core/server"
	"taskTracker/internal/adapters/ingress/routes"
	"taskTracker/internal/applications/tasktrackerservice"
)

func main() {
	// Initialize the database wrapper to manage database connections and operations
	dbWrapper, err := relationaldatabase.NewDbWrapper()
	if err != nil {
		log.Fatal(err) // Log and terminate if there is an error initializing the database wrapper
	}

	// signal handling for graceful shutdown
	shutDown := make(chan os.Signal, 1)
	// Notify on interrupt or termination signals
	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	// Create a new adapter for the database service layer
	database := iniservice.NewAdapter(dbWrapper)

	// Initialize application services, passing in the database adapter
	api := tasktrackerservice.NewApplication(database)

	// Set up the router for handling HTTP requests and define task tracker routes
	router := routes.NewTaskTrackerRouter().SetTaskTrackerRoutes(api)

	// for using a custom server implementation with shutdown capabilities,
	server.New(router, constants.DevEnvironment, constants.ApplicationPort, shutDown).Run()
	fmt.Println("shutting down server due to received signal: ", <-shutDown)
}
