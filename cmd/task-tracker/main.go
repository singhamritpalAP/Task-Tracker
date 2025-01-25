package main

import (
	"log"
	"net/http"
	"taskTracker/constants"
	"taskTracker/internal/adapters/ingress/routes"
	"taskTracker/internal/applications/tasktrackerservice"
	"taskTracker/internal/ports/tasktraceregress"
)

func main() {

	var database tasktraceregress.DbPort // fetch database object

	// Set up signal handling for graceful shutdown
	// shutDown := make(chan os.Signal, 1)
	// initializing services
	api := tasktrackerservice.NewApplication(database)
	// initializing routers
	router := routes.NewTaskTrackerRouter().SetTaskTrackerRoutes(api)
	err := http.ListenAndServe(constants.ApplicationPort, router)
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}
	//server.New(router, constants.DevEnvironment, constants.ApplicationPort, shutDown).Run()
}
