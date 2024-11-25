package main

import (
	"log"

	"github.com/garrettladley/gossip-gloomers/internal/echo"
	"github.com/garrettladley/gossip-gloomers/internal/gossipgloomers"
	"github.com/garrettladley/gossip-gloomers/internal/uniqueids"
)

func main() {
	if err := gossipgloomers.Run(echo.New, uniqueids.New); err != nil {
		log.Fatal(err)
	}
}
