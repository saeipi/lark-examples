package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *authService) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	resp = &pb_auth.SignUpResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}}
	var (
		u      = new(po.User)
		avatar *po.Avatar
	)
	u.Uid = xsnowflake.NewSnowflakeID()
	u.Password = req.Password
	u.Udid = req.Udid
	u.Status = 0
	u.Nickname = req.Nickname
	u.Firstname = req.Firstname
	u.Lastname = req.Lastname
	u.Gender = int(req.Gender)
	u.BirthTs = req.BirthTs
	u.Email = req.Email
	u.Mobile = req.Mobile
	u.RegPlatform = int(req.RegPlatform)
	//0:1~9
	u.ServerId = 10000
	u.CityId = int(req.CityId)
	u.AvatarKey = constant.CONST_AVATAR_KEY_SMALL
	avatar = &po.Avatar{
		GormEntityTs: entity.GormEntityTs{},
		OwnerId:      u.Uid,
		OwnerType:    1,
		AvatarSmall:  constant.CONST_AVATAR_KEY_SMALL,
		AvatarMedium: constant.CONST_AVATAR_KEY_MEDIUM,
		AvatarLarge:  constant.CONST_AVATAR_KEY_LARGE,
	}
	err = xmysql.Transaction(func(tx *gorm.DB) (terr error) {
		terr = s.authRepo.TxCreate(tx, u)
		if terr != nil {
			return
		}
		terr = s.avatarRepo.TxCreate(tx, avatar)
		if terr != nil {
			return
		}
		return
	})
	if err != nil {
		resp.Set(1001, "数据录入失败")
		return
	}
	resp.AccessToken, resp.RefreshToken, err = s.CreateToken(u)
	if err != nil {
		resp.AccessToken = nil
		resp.RefreshToken = nil
		resp.Set(1002, "生成token失败")
		return
	}
	copier.Copy(resp.UserInfo, u)
	copier.Copy(resp.UserInfo.Avatar, avatar)
	return
}

func (s *authService) CreateToken(u *po.User) (at *pb_auth.Token, rt *pb_auth.Token, err error) {
	var (
		accessToken  *xjwt.JwtToken
		refreshToken *xjwt.JwtToken
		maps         = map[string]string{}
		key          string
	)
	accessToken, err = xjwt.CreateToken(u.Uid, int32(u.RegPlatform), true, 7*24*3600)
	if err != nil {
		return
	}
	refreshToken, err = xjwt.CreateToken(u.Uid, int32(u.RegPlatform), true, 30*24*3600)
	if err != nil {
		return
	}
	at = &pb_auth.Token{
		Token:  accessToken.Token,
		Expire: accessToken.Expire,
	}
	rt = &pb_auth.Token{
		Token:  refreshToken.Token,
		Expire: refreshToken.Expire,
	}
	key = utils.ToString(u.Uid) + ":" + utils.ToString(u.RegPlatform)
	// TODO:缓存到redis
	maps[s.conf.Redis.Prefix+constant.RK_USER_ACCESS_TOKEN_SESSION_ID+key] = accessToken.SessionId
	maps[s.conf.Redis.Prefix+constant.RK_USER_REFRESH_TOKEN_SESSION_ID+key] = refreshToken.SessionId
	err = xredis.MSet(maps)
	return
}
