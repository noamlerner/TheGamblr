package server

import (
	"fmt"
	"net"

	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"
	"google.golang.org/grpc"
)

func RunCasino() {
	gs := grpc.NewServer()
	proto.RegisterCasinoServer(gs, NewCasinoServer())
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
