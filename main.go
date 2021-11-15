package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "go-micro-demo/proto"
)

type MeetServiceHandler struct {

}

func (m *MeetServiceHandler) Hello(ctx context.Context, in *proto.HelloRequest, out *proto.HelloResponse) error {
	out.Meeting = " 你好，" + in.Name
	return nil
}

func main()  {
	// 创建新的服务
	service := micro.NewService(micro.Name("Meet"))

	// 初始化，会解析命令行参数
	service.Init()

	// 注册处理器，调用 Meet 服务接口处理请求
	proto.RegisterMeetHandler(service.Server(), new(MeetServiceHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println("启动出错：", err)
	}
}