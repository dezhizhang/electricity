package main

import (
	"context"
	"electricity/proto"
	"electricity/utils"
	"fmt"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

//测试用户列表

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range rsp.Data {
		fmt.Println(user.Phone, user.NickName, user.Password, user.Birthday)
	}
}

// 测试创建用户

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp,err := userClient.CreateUser(context.Background(),&proto.CreateUserInfo{
			NickName: fmt.Sprintf("刘德华%d", i),
			Phone:    fmt.Sprintf("1599247844%d", i),
			Gender:   fmt.Sprintf("男"),
			Birthday: int64(1606832965),
			Password: utils.Md5String(string("123456")),
		})

		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)

	}
}

func main() {
	Init()
	TestCreateUser()
	conn.Close()
}
