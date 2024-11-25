package uniqueids

import (
	"strconv"
	"sync/atomic"

	"github.com/garrettladley/gossip-gloomers/internal/message"
	go_json "github.com/goccy/go-json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type UniqueIDs struct {
	message.Message
	ID string `json:"id"`
}

func New(n *maelstrom.Node) {
	var id atomic.Int64
	id.Add(1) // start at 1

	n.Handle("generate", func(msg maelstrom.Message) error {
		var _msg message.Message
		if err := go_json.Unmarshal(msg.Body, &_msg); err != nil {
			return err
		}

		var uids UniqueIDs
		uids.Message = _msg
		uids.Message.Type = message.GenerateOk
		uids.ID = n.ID() + "-" + strconv.Itoa(int(id.Load()))
		id.Add(1)

		return n.Reply(msg, uids)
	})
}
