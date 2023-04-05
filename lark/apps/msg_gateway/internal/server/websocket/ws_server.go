package websocket

import (
	"lark/apps/msg_gateway/internal/config"
	"lark/pkg/common/xgin"
	"lark/pkg/middleware"
)

type WsServer struct {
	hub       *Hub
	conf      *config.Config
	ginServer *xgin.GinServer
}

func NewWsServer(conf *config.Config) *WsServer {
	s := &WsServer{}
	s.hub = newHub()
	s.conf = conf
	s.ginServer = xgin.NewGinServer()
	s.addRouter()
	return s
}

func (s *WsServer) addRouter() {
	s.ginServer.Use(middleware.JwtAuth())
	s.ginServer.Engine.GET("ws", s.hub.Upgrade)
}

func (s *WsServer) Run() {
	go s.hub.run()
	s.ginServer.Run(s.conf.WsServer.Port)
}
