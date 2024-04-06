package main

import (
	"os"

	"github.com/snowmerak/regulus/regulus"
)

func main() {
	port := ""
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	if port == "" {
		port = ":8080"
	}

	server := regulus.New(port)

	println(regulus.AsciiArt)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
