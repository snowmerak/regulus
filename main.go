package main

import (
	"os"

	"github.com/snowmerak/regulus/regulus"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: regulus <port>")
	}

	port := os.Args[1]
	if port == "" {
		port = ":8080"
	}

	server := regulus.New(port)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
