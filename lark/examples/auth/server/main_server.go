package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"lark/pkg/proto/pb_auth"
	"net"
)

func main() {
	var s = new(Server)
	s.Run()
}

type Server struct {
	pb_auth.UnimplementedAuthServer
}

func (s *Server) Run() {
	var (
		listener net.Listener
		err      error
		srv      *grpc.Server
	)
	listener, err = net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	srv = grpc.NewServer()
	pb_auth.RegisterAuthServer(srv, s)
	if err = srv.Serve(listener); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}

func (s *Server) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	resp = new(pb_auth.SignUpResp)
	resp.Msg = "注册成功"
	return
}
