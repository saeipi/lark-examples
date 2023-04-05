package svc_auth

import (
	"github.com/jinzhu/copier"
	auth_client "lark/apps/auth/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

type AuthService interface {
	SignUp(req *dto_auth.SignUpReq) (resp *xhttp.Resp)
}

type authService struct {
	conf       *config.Config
	authClient auth_client.AuthClient
}

func NewAuthService(conf *config.Config) AuthService {
	authClient := auth_client.NewAuthClient(conf.Etcd, conf.AuthServer, conf.Jaeger, conf.Name)
	return &authService{conf: conf, authClient: authClient}
}

func (s *authService) SignUp(req *dto_auth.SignUpReq) (resp *xhttp.Resp) {
	var (
		r = new(pb_auth.SignUpReq)
	)
	resp = new(xhttp.Resp)
	copier.Copy(r, req)

	var resp1 = s.authClient.SignUp(r)
	if resp1 == nil {
		//服务故障
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if resp1.Code > 0 {
		//业务逻辑报错
		resp.SetResult(resp.Code, resp.Msg)
		return
	}
	resp.Data = &dto_auth.SignUpResp{
		UserInfo:     resp1.UserInfo,
		AccessToken:  resp1.AccessToken,
		RefreshToken: resp1.RefreshToken,
	}
	return
}
