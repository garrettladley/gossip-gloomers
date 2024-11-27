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

	n.Handle("generate", func(req maelstrom.Message) error {
		var msg message.Message
		if err := go_json.Unmarshal(req.Body, &msg); err != nil {
			return err
		}

		var uids UniqueIDs
		uids.Message = msg
		uids.Message.Type = message.GenerateOk
		uids.ID = n.ID() + "-" + strconv.Itoa(int(id.Load()))
		id.Add(1)

		return n.Reply(req, uids)
	})
}
