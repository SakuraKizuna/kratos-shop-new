package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	v1 "user/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()

	TestCreateUser() // 创建用户

	conn.Close()
}

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(any("grpc link err" + err.Error()))
	}
	userClient = v1.NewUserClient(conn)
}

func TestCreateUser() {

	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserInfo{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		Password: "admin123",
		NickName: fmt.Sprintf("YWWW%d", 1),
	})
	if err != nil {
		spew.Dump(err)
	}
	fmt.Println(rsp)

	rsp2, err2 := userClient.GetUserInfo(context.Background(), &v1.GetUserRequest{
		Id: int64(3),
	})
	if err2 != nil {
		spew.Dump("err2:", err2)
	}
	fmt.Println("rsp2:", rsp2)

}
