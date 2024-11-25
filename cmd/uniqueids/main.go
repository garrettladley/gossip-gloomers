package main

import (
	"log"

	"github.com/garrettladley/gossip-gloomers/internal/echo"
	"github.com/garrettladley/gossip-gloomers/internal/gossipgloomers"
	"github.com/garrettladley/gossip-gloomers/internal/uniqueids"
)

func main() {
	log.Fatal(gossipgloomers.Run(echo.New, uniqueids.New))
}
