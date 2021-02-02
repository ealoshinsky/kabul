package hub

import (
	"log"

	core "github.com/AsynkronIT/protoactor-go/actor"
)

type hub struct{}

func (state *hub) Receive() {

}

//NewHub create new group of workers
func NewHub() {
	// only e.g stub
	system := core.NewActorSystem()
	log.Println(system)
}
