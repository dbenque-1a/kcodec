package main

import "github.com/dbenque/kcodec/pkg/server"

func main() {
	server := &server.Server{
		BindAddress: ":8080",
	}

	server.Init()
	server.Run()
}
