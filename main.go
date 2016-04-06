package main

import (
	"flag"
	"log"

	"github.com/alyoshka/calc-server/server"
)

var host = flag.String("host", "localhost:8080", "Host and port to listen")

func main() {
	flag.Parse()
	if err := server.Run(*host); err != nil {
		log.Panic("error running server:", err)
	}
}
