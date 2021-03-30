package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	config "grpcApply/conf"
	"grpcApply/protos"
	"grpcApply/server"
	"log"
	"net"
)

func main() {
	config.Load()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Conf.Tcp.Address, config.Conf.Tcp.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	registered(lis)
}

func registered(lis net.Listener) {
	es := &server.EchoServer{}

	gs := grpc.NewServer()

	protos.RegisterEchoServer(gs, es)
	reflection.Register(gs)
	log.Printf("GRPC TCP Server run on %s:%s", config.Conf.Tcp.Address, config.Conf.Tcp.Port)

	if err := gs.Serve(lis); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}
}
