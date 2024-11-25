package gossipgloomers

import (
	"log"

	"github.com/garrettladley/gossip-gloomers/internal/handler"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Run(h ...handler.Handler) {
	n := maelstrom.NewNode()
	for _, handler := range h {
		handler(n)
	}
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
