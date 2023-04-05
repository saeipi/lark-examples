package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(gin.Recovery())
	return &GinServer{Engine: engine}
}

func (s *GinServer) Run(port int) {
	var (
		addr string
		err  error
	)
	addr = ":" + strconv.Itoa(port)
	err = s.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed:", err.Error())
	}
}

func (s *GinServer) Use(m ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.Use(m...)
}
