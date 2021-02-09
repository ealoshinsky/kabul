package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ealoshinsky/kabul/hub"
	"github.com/ealoshinsky/kabul/lib"
	monitor "github.com/ealoshinsky/kabul/monitor"
	log "github.com/sirupsen/logrus"
)

var configPath string // configuration file path

func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "The configuration file path.")
}

func main() {
	flag.Parse() // parse cmd arguments

	var config = &lib.Config{}

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.WithFields(log.Fields{
			"parse-config": "error-parse",
			"detail":       err.Error(),
		}).Fatal("Failure parse configuration file.")
		os.Exit(1)
	}

	log.Println(config)

	go monitor.StartMonitor()
	go hub.NewHub(config)
	for {
	}
}
