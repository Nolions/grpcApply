package server

import (
	"context"
	"grpcApply/protos"
	"log"
)

type EchoServer struct{}

func (e *EchoServer) SayHello(ctx context.Context, req *protos.EchoRequest) (resp *protos.EchoReply, err error) {
	log.Printf("receive client request, client send:%s\n", req.Message)
	return &protos.EchoReply{
		Message: req.Message,
	}, nil
}
