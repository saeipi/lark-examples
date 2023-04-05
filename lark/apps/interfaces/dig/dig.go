package dig

import (
	"go.uber.org/dig"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"
	"lark/apps/interfaces/internal/service/svc_auth"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	//container.Provide(server.NewServer)
	container.Provide(ctrl_auth.NewAuthCtrl)
	container.Provide(svc_auth.NewAuthService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
