package main

import (
	"log"

	"github.com/alyoshka/calc-server/server"
)

func main() {
	if err := server.Run("localhost:9090"); err != nil {
		log.Panic("error running server:", err)
	}
}
