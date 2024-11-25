package main

import (
	"log"

	"github.com/garrettladley/gossip-gloomers/internal/echo"
	"github.com/garrettladley/gossip-gloomers/internal/gossipgloomers"
)

func main() {
	log.Fatal(gossipgloomers.Run(echo.Register))
}
