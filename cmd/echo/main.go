package main

import (
	"github.com/garrettladley/gossip-gloomers/internal/echo"
	"github.com/garrettladley/gossip-gloomers/internal/gossipgloomers"
)

func main() {
	gossipgloomers.Run(echo.Register)
}
