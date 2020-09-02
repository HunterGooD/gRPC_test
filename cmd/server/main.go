package main

import (
	"log"
	"net"

	"github.com/HunterGooD/gRPC_test/pkg/adder"
	"github.com/HunterGooD/gRPC_test/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	serv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, serv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
