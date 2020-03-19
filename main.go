// main package, app entry point
package main

import (
	"flag"
	"fmt"
	"go-api-boilerplate/lib/config"
	"go-api-boilerplate/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	config.Init(*environment)

	server.Init()
}
