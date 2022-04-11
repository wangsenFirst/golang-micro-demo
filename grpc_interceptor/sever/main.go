package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net"

	"google.golang.org/grpc"
	"grpc_demo/proto/proto"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *helloword.HelloRequest) (*helloword.HelloReply, error) {
	// 从metadate取数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	return &helloword.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	// 实现拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到一个请求。。。")
		i, err := handler(ctx, req)
		fmt.Println("请求结束。。。")
		return i, err
	}
	unaryInterceptor := grpc.UnaryInterceptor(interceptor)

	// 创建grpc server
	g := grpc.NewServer(unaryInterceptor)

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
