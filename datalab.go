package datalab

import (
	_ "embed"

	"github.com/gabereiser/datalab/config"
	"github.com/gabereiser/datalab/log"
	"github.com/gabereiser/datalab/network"
)

//go:generate bash .ops/version.sh
//go:embed version.txt
var Version string

var Banner string = `
 ____        _        _          _     
|  _ \  __ _| |_ __ _| |    __ _| |__  
| | | |/ _' | __/ _' | |   / _' | '_ \ 
| |_| | (_| | || (_| | |__| (_| | |_) |
|____/ \__,_|\__\__,_|_____\__,_|_.__/ 
`

func Run() {
	log.Info("secret key is %s", config.Config.SecretKey)
	server := network.NewWebServer()
	server.Listen(config.Config.Domain + ":8080")
}
