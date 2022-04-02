package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_demo/proto/proto"
)

func main() {
	// grpc拨号
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := helloword.NewGreeterClient(conn)
	// 往matedate 里面放数据
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":    "sen",
		"pasword": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &helloword.HelloRequest{Name: "sen"})

	// r, err := c.SayHello(context.Background(), &helloword.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
