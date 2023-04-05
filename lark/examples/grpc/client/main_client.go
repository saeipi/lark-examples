package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lark/pkg/proto/pb_auth"
)

func main() {
	req := &pb_auth.SignUpReq{
		Nickname: "lark",
	}
	conn, err := grpc.Dial("127.0.0.1:6600", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	client := pb_auth.NewAuthClient(conn)
	var resp *pb_auth.SignUpResp
	resp, err = client.SignUp(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp == nil {
		return
	}
	fmt.Println(resp.Msg)

}
