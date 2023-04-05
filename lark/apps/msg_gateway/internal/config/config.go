package config

import (
	"flag"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name     string         `yaml:"name"`
	ServerID int            `yaml:"server_id"`
	Log      string         `yaml:"log"`
	WsServer *conf.WsServer `yaml:"ws_server"`
	Redis    *conf.Redis    `yaml:"redis"`
}

var (
	config = new(Config)
)

var (
	confFile    = flag.String("cfg", "./configs/msg_gateway.yaml", "config file")
	grpcPort    = flag.Int("gp", 7300, "grpc server port")
	wsPort      = flag.Int("wp", 7301, "websocker server port")
	monitorPort = flag.Int("mp", 7302, "metrics server port")
	serverId    = flag.Int("sid", 1, "server id")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId

	config.WsServer.Port = *wsPort
	config.WsServer.ServerId = config.ServerID

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+utils.IntToStr(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
