package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_demo/proto/proto"
	"time"
)

func main() {
	// 拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)

		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}
	unaryInterceptor := grpc.WithUnaryInterceptor(interceptor)
	// grpc拨号
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), unaryInterceptor)
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
