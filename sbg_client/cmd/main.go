package main

import (
	"sbg_client/internal/webpage"

	"github.com/alexflint/go-arg"

	"sbg_server/pkg/logtools"

	log "github.com/sirupsen/logrus"
)

func main() {
	var args struct {
		logtools.LogCfg
		webpage.RouterCfg
	}
	arg.MustParse(&args)

	logtools.InitLog(args.LogCfg)

	log.Info("Client is starting...")

	log.Infof("Start arguments: %+v", args)

	log.Print("HTTP server is starting...")

	router, err := webpage.NewRouter(args.RouterCfg)
	if err != nil {
		logtools.LogErrorWithStack(err).Fatal("Can`t create router")
	}

	if err := router.Run(); err != nil {
		log.WithError(err).Fatal("Can`t start HTTP server")
	}
}
