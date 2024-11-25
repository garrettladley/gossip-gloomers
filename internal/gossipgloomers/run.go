package gossipgloomers

import (
	"github.com/garrettladley/gossip-gloomers/internal/handler"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Run(h ...handler.Handler) error {
	n := maelstrom.NewNode()
	for _, handler := range h {
		handler(n)
	}
	return n.Run()
}
