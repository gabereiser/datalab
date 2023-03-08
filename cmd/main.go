package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabereiser/datalab/log"

	"github.com/gabereiser/datalab"
)

var sigterm chan os.Signal = make(chan os.Signal)

func main() {
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println(datalab.Banner)
	log.Info("DATALAB %s", datalab.Version)
	log.Info("Starting Datalab")
	datalab.Run()
	<-sigterm
	fmt.Println()
	log.Info("Stopping lab...")
}
