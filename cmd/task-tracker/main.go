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

	dbWrapper, err := relationaldatabase.NewDbWrapper()
	if err != nil {
		log.Fatal(err)
	}
	database := iniservice.NewAdapter(dbWrapper)
	// initializing services
	api := tasktrackerservice.NewApplication(database)
	// initializing routers
	router := routes.NewTaskTrackerRouter().SetTaskTrackerRoutes(api)
	err = http.ListenAndServe(constants.ApplicationPort, router)
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}
	//server.New(router, constants.DevEnvironment, constants.ApplicationPort, shutDown).Run()
}
