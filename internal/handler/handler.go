package handler

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Handler func(n *maelstrom.Node)
