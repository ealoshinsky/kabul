package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ealoshinsky/kabul/hub"
	monitor "github.com/ealoshinsky/kabul/monitor"
	log "github.com/sirupsen/logrus"
)

var configPath string // configuration file path

// Config specificaton for all in application
type Config struct {
	Monitor struct {
		Port string
		Addr string
	} `toml:"monitor"`
}

func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "The configuration file path.")
}

func main() {
	flag.Parse() // parse cmd arguments

	var config = &Config{}

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.WithFields(log.Fields{
			"parse-data": "error-parse",
			"detail":     err.Error(),
		}).Fatal("Failure parse configuration file.")
		os.Exit(-1)
	}

	go monitor.StartMonitor()
	go hub.NewHub()
	for {
	}
}
