package echo

import (
	go_json "github.com/goccy/go-json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type MessageType string

const (
	Echo   MessageType = "echo"
	EchoOk MessageType = "echo_ok"
)

type EchoRequest struct {
	Type      MessageType `json:"type"`
	MessageID uint        `json:"msg_id"`
	Echo      string      `json:"echo"`
}

func Register(n *maelstrom.Node) {
	n.Handle("echo", func(msg maelstrom.Message) error {
		var er EchoRequest
		if err := go_json.Unmarshal(msg.Body, &er); err != nil {
			return err
		}

		er.Type = EchoOk

		return n.Reply(msg, er)
	})
}
