package server

import "lark/apps/auth/internal/server/auth"

type Server struct {
	authServer auth.AuthServer
}

func NewServer(authServer auth.AuthServer) *Server {
	return &Server{authServer: authServer}
}

func (s *Server) Run() {
	s.authServer.Run()
}
