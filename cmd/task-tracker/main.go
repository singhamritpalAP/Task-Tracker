package main

import (
	"log"
	"net/http"
	"taskTracker/constants"
	"taskTracker/core/iniservice"
	"taskTracker/core/relationaldatabase"
	"taskTracker/internal/adapters/ingress/routes"
	"taskTracker/internal/applications/tasktrackerservice"
)

func main() {
	// Initialize the database wrapper to manage database connections and operations
	dbWrapper, err := relationaldatabase.NewDbWrapper()
	if err != nil {
		log.Fatal(err) // Log and terminate if there is an error initializing the database wrapper
	}

	// Create a new adapter for the database service layer
	database := iniservice.NewAdapter(dbWrapper)

	// Initialize application services, passing in the database adapter
	api := tasktrackerservice.NewApplication(database)

	// Set up the router for handling HTTP requests and define task tracker routes
	router := routes.NewTaskTrackerRouter().SetTaskTrackerRoutes(api)

	// Start the HTTP server on the specified port, using the defined router for handling requests
	err = http.ListenAndServe(constants.ApplicationPort, router)
	if err != nil {
		log.Fatal("Error starting server", err) // Log any errors that occur while starting the server
		return
	}

	// for using a custom server implementation with shutdown capabilities,
	// minor changes needed in custom server logic to use custom server
	// server.New(router, constants.DevEnvironment, constants.ApplicationPort, shutDown).Run()
}
