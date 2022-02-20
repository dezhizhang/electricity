package main

import (
	"electricity/proto"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	conn,err := grpc.Dial("127.0.0.1:8000",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient =  proto.NewUserClient(conn)
}

func main()  {
	Init()
	conn.Close()
}

