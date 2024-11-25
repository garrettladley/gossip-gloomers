package main

import (
	"log"

	"github.com/garrettladley/gossip-gloomers/internal/echo"
	"github.com/garrettladley/gossip-gloomers/internal/gossipgloomers"
)

func main() {
	if err := gossipgloomers.Run(echo.New); err != nil {
		log.Fatal(err)
	}
}
