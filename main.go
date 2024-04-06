package main

import "github.com/snowmerak/regulus/regulus"

func main() {
	server := regulus.New(":8080")

	if err := server.Start(); err != nil {
		panic(err)
	}
}
