package main

import (
	"github.com/alexflint/go-arg"

	log "github.com/sirupsen/logrus"

	"sbg_server/pkg/logtools"
)

type Args struct {
	logtools.LogCfg
}

func main() {
	var args Args
	arg.MustParse(&args)

	logtools.InitLog(args.LogCfg)

	log.Info("SBG server is starting...")

	log.Infof("Start arguments: %+v", args)

}
