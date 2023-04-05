package main

import (
	"lark/apps/auth/dig"
	"lark/apps/auth/internal/config"
	"lark/apps/auth/internal/server"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
	"sync"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
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
