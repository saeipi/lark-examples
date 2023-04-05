package main

import (
	"lark/apps/msg_gateway/dig"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server"
	"lark/pkg/common/xredis"
	"sync"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	wg := sync.WaitGroup{}
	dig.Invoke(func(srv *server.Server) {
		srv.Run()
	})
	wg.Add(1)
	wg.Wait()
}
