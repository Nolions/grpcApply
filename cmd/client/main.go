package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	config "grpcApply/conf"
	"grpcApply/protos"
	"log"
)

func main() {
	config.Load()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Conf.Tcp.Address, config.Conf.Tcp.Port), grpc.WithInsecure())
	log.Printf("Connection to %s:%s", config.Conf.Tcp.Address, config.Conf.Tcp.Port)
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := protos.NewEchoClient(conn)
	hello(c)
}

func hello(c protos.EchoClient) {
	r, err := c.SayHello(context.Background(), &protos.EchoRequest{Message: "HI HI HI HI"})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s", r.Message)
}
