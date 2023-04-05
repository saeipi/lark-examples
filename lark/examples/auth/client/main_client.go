package main

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lark/examples/auth/constant"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
)

func main() {
	var (
		target = constant.GRPC_SERVER_ADDR
		conn   *grpc.ClientConn
		err    error
		client pb_auth.AuthClient
		req    *pb_auth.SignUpReq
		resp   *pb_auth.SignUpResp
	)

	req = &pb_auth.SignUpReq{
		RegPlatform: pb_enum.PLATFORM_TYPE_WINDOWS,
		Nickname:    "lark",
		Password:    "lark2023",
		Firstname:   "H",
		Lastname:    "CJ",
		Gender:      1,
		BirthTs:     1676188896,
		Email:       "ksert@163.com",
		Mobile:      "18111111111",
		AvatarKey:   "",
		CityId:      1,
		Code:        1234,
		Udid:        uuid.NewV4().String(),
		ServerId:    1,
	}
	conn, err = grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("创建客户端连接错误:", err.Error())
		return
	}
	defer conn.Close()
	client = pb_auth.NewAuthClient(conn)
	resp, err = client.SignUp(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp == nil {
		fmt.Println("服务端故障")
		return
	}
	fmt.Println(resp.Msg)
}
