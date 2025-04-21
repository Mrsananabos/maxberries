package main

import (
	"reviewsService/http/rest"
)

func main() {
	server, err := rest.NewServer()
	if err != nil {
		panic(err)
	}

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
