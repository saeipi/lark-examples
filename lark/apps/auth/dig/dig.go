package dig

import (
	"go.uber.org/dig"
	"lark/apps/auth/internal/config"
	"lark/apps/auth/internal/server"
	"lark/apps/auth/internal/server/auth"
	"lark/apps/auth/internal/service"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(auth.NewAuthServer)
	container.Provide(service.NewAuthService)
	container.Provide(repo.NewAuthRepository)
	container.Provide(repo.NewAvatarRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
