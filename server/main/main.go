package main

import (
	"fmt"
	"net"

	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"
	"github.com/noamlerner/TheGamblr/server"

	"google.golang.org/grpc"
)

func main() {
	gs := grpc.NewServer()
	proto.RegisterCasinoServer(gs, server.NewCasinoServer())
	fmt.Printf("starting server at port 8080")
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	err = gs.Serve(listen)
	if err != nil {
		panic(err)
	}
}
