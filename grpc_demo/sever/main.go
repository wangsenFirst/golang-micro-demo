package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"grpc_demo/proto/proto"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *helloword.HelloRequest) (*helloword.HelloReply, error) {
	return &helloword.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	fmt.Println("66666")
	// 创建grpc server
	g := grpc.NewServer()

	s := Server{}
	// 注册写好的grpc service
	helloword.RegisterGreeterServer(g, &s)
	// 创建监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":8080"))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	// 启动grpc监听
	g.Serve(lis)
}
