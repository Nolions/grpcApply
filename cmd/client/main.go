package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcApply/protos"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:4321", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := protos.NewEchoClient(conn)
	hello(c)
}

func hello(c protos.EchoClient)  {
	r, err := c.SayHello(context.Background(), &protos.EchoRequest{Message: "HI HI HI HI"})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s", r.Message)
}
