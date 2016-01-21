package main

import (
	"./server"
	"flag"
	"log"
)

var (
	flListenAddr = flag.String("l", "127.0.0.1:17017", "http listen addr:port")
)

func main() {
	flag.Parse()
	app := server.NewApp()
	if err := app.Run(*flListenAddr); err != nil {
		log.Fatalf("run error: %s", err)
	}
}
