package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gabereiser/datalab/log"

	"github.com/gabereiser/datalab"
)

var sigterm chan os.Signal = make(chan os.Signal)

func main() {
	signal.Notify(sigterm, os.Interrupt)
	fmt.Println(datalab.Banner)
	log.Info("DATALAB %s", datalab.Version)
	log.Info("Starting Datalab")
	go func() {
		<-sigterm
		log.Info("Shutting down...")
		os.Exit(0)
	}()
	defer datalab.Shutdown()
	datalab.Run()
}
