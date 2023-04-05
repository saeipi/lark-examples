package service

import (
	"context"
	"lark/apps/auth/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_auth"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error)
	SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
	SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error)
}

type authService struct {
	authRepo   repo.AuthRepository
	avatarRepo repo.AvatarRepository
	conf       *config.Config
}

func NewAuthService(authRepo repo.AuthRepository, avatarRepo repo.AvatarRepository, conf *config.Config) AuthService {
	return &authService{authRepo: authRepo, avatarRepo: avatarRepo, conf: conf}
}

func (s *authService) SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error) {
	return
}

func (s *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return
}

func (s *authService) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error) {
	return
}
