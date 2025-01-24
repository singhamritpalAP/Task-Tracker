package main

import (
	"os"
	"taskTracker/constants"
	"taskTracker/core/server"
	"taskTracker/internal/adapters/ingress/routes"
	"taskTracker/internal/applications/tasktrackerservice"
	"taskTracker/internal/ports/tasktraceregress"
)

func main() {

	var database tasktraceregress.DbPort // fetch database object

	// Set up signal handling for graceful shutdown
	shutDown := make(chan os.Signal, 1)
	// initializing services
	api := tasktrackerservice.NewApplication(database)
	// initializing routers
	router := routes.NewTaskTrackerRouter().SetTaskTrackerRoutes(api)
	server.New(router, constants.DevEnvironment, constants.ApplicationPort, shutDown).Run()
}
