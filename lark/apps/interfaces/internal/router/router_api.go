package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	openGroup := engine.Group("open")
	registerOpenRouter(openGroup)

	apiGroup := engine.Group("api")
	apiGroup.Use(middleware.JwtAuth())
	registerApiRouter(apiGroup)
}

func registerOpenRouter(group *gin.RouterGroup) {
	authGroup := group.Group("auth")

	var ctrl *ctrl_auth.AuthCtrl
	dig.Invoke(func(c *ctrl_auth.AuthCtrl) {
		ctrl = c
	})
	authGroup.POST("sign_in", ctrl.SignIn)
	authGroup.POST("sign_up", ctrl.SignUp)
}

func registerApiRouter(group *gin.RouterGroup) {
	authGroup := group.Group("auth")
	var ctrl *ctrl_auth.AuthCtrl
	dig.Invoke(func(c *ctrl_auth.AuthCtrl) {
		ctrl = c
	})
	authGroup.POST("sign_out", ctrl.SignOut)
}
