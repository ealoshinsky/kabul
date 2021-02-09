package hub

import (
	"os"

	log "github.com/sirupsen/logrus"

	core "github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ealoshinsky/kabul/lib"
)

type hub struct{}

func (state *hub) Receive() {

}

//NewHub create new group of workers
func NewHub(config *lib.Config) {

	system := core.NewActorSystem()
	if len(config.Hub) <= 0 {
		log.WithFields(log.Fields{
			"parse-config": "error-value-hub",
			"detail":       "No defined hub workers",
		}).Fatal("Could not start hub subsystem.")
		os.Exit(0)
	}

	println(system)
}
