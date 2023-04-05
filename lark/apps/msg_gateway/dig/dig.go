package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server"
	"lark/apps/msg_gateway/internal/server/websocket"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(websocket.NewWsServer)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
