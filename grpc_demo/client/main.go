package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	r, err := c.SayHello(context.Background(), &helloword.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
