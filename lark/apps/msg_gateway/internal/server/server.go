package server

import "lark/apps/msg_gateway/internal/server/websocket"

type Server struct {
	ws *websocket.WsServer
}

func NewServer(ws *websocket.WsServer) *Server {
	return &Server{ws: ws}
}

func (s *Server) Run() {
	s.ws.Run()
}
