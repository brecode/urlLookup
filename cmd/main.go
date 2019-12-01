package main

import (
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/brecode/urlLookup/handler"
	"github.com/brecode/urlLookup/router"
)

const (
	defaultAddress = "0.0.0.0"
	defaultPort    = "33333"
)

func main() {
	var err error
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	// create a new handler initialize and inject any dependencies
	handler := handler.NewHandler(handler.UseHandlerDeps(
		func(deps *handler.Deps) {
			deps.Logger = logger
		}))

	// create a new router initialize and inject any dependencies
	rtr := router.NewRouter(router.UseRouterDeps(
		func(deps *router.Deps) {
			deps.Handler = handler
			deps.Logger = logger
		}))
	if err = rtr.Init(); err != nil {
		logger.Fatalf("Could not initialize router, exiting with error: %+v", err)
	}

	// go thread to handle interrupt signal
	go handleInterrupt(logger)

	logger.Printf("Starting HTTP service at >> %s:%s\n", defaultAddress, defaultPort)
	logger.Fatal(http.ListenAndServe(defaultAddress+":"+defaultPort, rtr))
}

func handleInterrupt(logger *log.Logger) {
	// handle interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	for {
		select {
		case <-quit:
			logger.Debugf("[MAIN]: Caught interrupt signal, exiting...")
			os.Exit(0)
		}
	}
}
