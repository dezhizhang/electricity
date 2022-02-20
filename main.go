package main

import (
	"electricity/proto"
	"electricity/service"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	IP := flag.String("ip","127.0.0.1","ip地址")
	PORT := flag.Int("port",8000,"端口")

	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server,&service.UserServer{})
	listen,err := net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*PORT))
	if err != nil {
		panic("fail to listen:" + err.Error())
	}
	err = server.Serve(listen)
	if err != nil  {
		panic("fail to start grpc:" + err.Error())
	}
}
