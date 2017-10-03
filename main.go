package main

import (
	"flag"
	"fmt"
	"os"

	"./lib/config"
	"./server"
)

func main() {
	enviroment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	config.Init(*enviroment)

	server.Init()
}
