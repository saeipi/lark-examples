package ctrl_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/apps/interfaces/internal/service/svc_auth"
	"lark/pkg/xhttp"
)

type AuthCtrl struct {
	authService svc_auth.AuthService
}

func NewAuthCtrl(authService svc_auth.AuthService) *AuthCtrl {
	return &AuthCtrl{authService: authService}
}

func (ctrl AuthCtrl) SignIn(ctx *gin.Context) {
	fmt.Println("访问了SignIn Api")
	ctx.SecureJSON(200, "调用api成功")
}

func (ctrl AuthCtrl) SignUp(ctx *gin.Context) {
	var (
		req  dto_auth.SignUpReq
		resp *xhttp.Resp
		err  error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR)
		return
	}
	resp = ctrl.authService.SignUp(&req)
	if resp == nil {
		//连接grpc服务失败
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		//业务错误
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

func (ctrl AuthCtrl) SignOut(ctx *gin.Context) {
	fmt.Println("SignOut Api")
	ctx.SecureJSON(200, "调用api成功")
}
