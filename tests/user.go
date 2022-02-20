package main

import (
	"context"
	"electricity/proto"
	"fmt"
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

func TestGetUserList()  {
	rsp,err := userClient.GetUserList(context.Background(),&proto.PageInfo{
		Page: 1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}

	for _,user := range rsp.Data {
		fmt.Println(user.Phone,user.NickName,user.Password,user.Birthday)
	}
}

func main()  {
	Init()
	conn.Close()
}

