package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcApply/protos"
	"grpcApply/server"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:4321")
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
	if err := gs.Serve(lis); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}
}
