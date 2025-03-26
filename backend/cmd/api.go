package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pecet3/las-test-pdf/cmd/repos"
	"github.com/pecet3/las-test-pdf/cmd/router"
	"github.com/pecet3/las-test-pdf/utils"
	"github.com/pecet3/logger"
)

func runAPI() {
	logger.Info("Starting...")
	utils.LoadEnv()
	app := repos.NewApp()
	address := os.Getenv("ADDRESS")
	router.Run(app)

	server := &http.Server{
		Addr:    address,
		Handler: app.Srv,
	}

	logger.Info(fmt.Sprintf("Server is listening on: [%s]", address))
	log.Fatal(server.ListenAndServe())

}
