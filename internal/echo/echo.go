package echo

import (
	"github.com/garrettladley/gossip-gloomers/internal/message"
	go_json "github.com/goccy/go-json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Echo struct {
	message.Message
	Echo string `json:"echo"`
}

func New(n *maelstrom.Node) {
	n.Handle("echo", func(msg maelstrom.Message) error {
		var echo Echo
		if err := go_json.Unmarshal(msg.Body, &echo); err != nil {
			return err
		}

		echo.Type = message.EchoOk

		return n.Reply(msg, echo)
	})
}
