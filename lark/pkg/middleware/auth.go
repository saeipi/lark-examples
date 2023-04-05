package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token     *jwt.Token
			maps      jwt.MapClaims
			uid       interface{}
			platform  interface{}
			sessionId interface{}
			sId       string
			key       string
			ok        bool
			err       error
		)
		//验证jwt
		token, err = xjwt.ParseFromCookie(ctx)
		if err != nil {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_JWT_TOKEN_ERR, xhttp.ERROR_HTTP_JWT_TOKEN_ERR)
			return
		}
		maps = token.Claims.(jwt.MapClaims)
		uid, ok = maps[constant.USER_UID]
		if ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
			return
		}
		platform, ok = maps[constant.USER_PLATFORM]
		if ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
			return
		}
		sessionId, ok = maps[constant.USER_JWT_SESSION_ID]
		if ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_SESSION_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_SESSION_ID_DOESNOT_EXIST)
			return
		}
		key = constant.RK_USER_ACCESS_TOKEN_SESSION_ID + utils.ToString(uid) + ":" + utils.ToString(platform)
		sId, err = xredis.Get(key)
		if err == redis.Nil {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_NOT_AUTHORIZED, xhttp.ERROR_HTTP_REQ_NOT_AUTHORIZED)
			return
		}
		if err != nil {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REDIS_GET_FAILED, xhttp.ERROR_HTTP_REDIS_GET_FAILED)
			return
		}
		if sId != utils.ToString(sessionId) {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_NOT_AUTHORIZED, xhttp.ERROR_HTTP_REQ_NOT_AUTHORIZED)
			return
		}
		ctx.Set(constant.USER_UID, uid)
		ctx.Set(constant.USER_PLATFORM, platform)
	}
}
