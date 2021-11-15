package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "go-micro-demo/proto"
)

func main() {
	// 创建一个新的服务
	service := micro.NewService(micro.Name("Meet.Client"))
	// 初始化
	service.Init()

	// 创建 Meet 客户端
	// 注意这里的 "Meet" 要和服务端 main.go 中的 micro.Name("Meet") 注册的 "Meet" 名称要一致，
	// 否则会报错
	meet := proto.NewMeetService("Meet", service.Client())

	// 远程调用 Meet 服务的 Hello 方法
	rsp, err := meet.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Alex",
	})

	if err != nil {
		fmt.Println("客户端调用报错：", err)
		return
	}

	// 打印返回结果
	fmt.Println("response ==> ", rsp.Meeting)

}
