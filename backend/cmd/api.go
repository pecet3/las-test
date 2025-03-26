package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pecet3/las-test-pdf/cmd/repos"
	"github.com/pecet3/las-test-pdf/cmd/router"
	"github.com/pecet3/las-test-pdf/utils"
	"github.com/pecet3/logger"
)

const BASE_URL = "0.0.0.0:9090"

func runAPI() {
	logger.Info("Starting...")
	utils.LoadEnv()
	app := repos.NewApp()

	router.Run(app)

	server := &http.Server{
		Addr:    BASE_URL,
		Handler: app.Srv,
	}

	logger.Info(fmt.Sprintf("Server is listening on: [%s]", BASE_URL))
	log.Fatal(server.ListenAndServe())

}
