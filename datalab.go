package datalab

import (
	"embed"
	_ "embed"

	"github.com/gabereiser/datalab/config"
	"github.com/gabereiser/datalab/data"
	"github.com/gabereiser/datalab/log"
	"github.com/gabereiser/datalab/network"
)

var Version string = "dev"

//go:generate bash .ops/vite.sh
//go:embed public/*
var viewsfs embed.FS

var Banner string = `
 ____        _        _          _     
|  _ \  __ _| |_ __ _| |    __ _| |__  
| | | |/ _' | __/ _' | |   / _' | '_ \ 
| |_| | (_| | || (_| | |__| (_| | |_) |
|____/ \__,_|\__\__,_|_____\__,_|_.__/ 
`

var server *network.WebServer

func Run() {
	log.Info("secret key is %s", config.Config.SecretKey)
	data.Migrate()
	server = network.NewWebServer(viewsfs)
	server.SetupRoutes()
	server.Listen(config.Config.Domain + ":8080")
}

func Shutdown() {
	server.Stop()
}
