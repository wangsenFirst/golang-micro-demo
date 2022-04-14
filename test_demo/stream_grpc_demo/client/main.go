package main

import (
	"google.golang.org/grpc"

	"context"
	_ "google.golang.org/grpc/balancer/grpclb"
	"log"
	"stream_grpc/proto"
	"time"
)

const (
	ADDRESS = "localhost:50052"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := proto.NewGreeterClient(conn)
	//调用服务端推送流
	reqstreamData := &proto.StreamReqData{Data: "调用服务端推送流"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		// 服务段结束返回会返回一个err是EOF
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&proto.StreamReqData{Data: "客户端 推送 流"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&proto.StreamReqData{Data: "服务端 客户端 双向流"})
			time.Sleep(time.Second)
		}
	}()

	select {}

}
