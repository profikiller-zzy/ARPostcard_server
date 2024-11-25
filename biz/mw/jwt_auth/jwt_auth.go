package jwt_auth

import (
	"context"
	"github.com/RanFeng/ierror"
	"github.com/cloudwego/hertz/pkg/app"
	"uav/biz/consts"
	"uav/biz/utils"
)

func JwtAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("token")
		if token == "" {
			utils.RespErr(ctx, c, ierror.NewIError(-1, "未携带token"))
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.RespErr(ctx, c, ierror.NewIError(-1, "token错误"))
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("token")
		if token == "" {
			utils.RespErr(ctx, c, ierror.NewIError(-1, "未携带token"))
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.RespErr(ctx, c, ierror.NewIError(-1, "token错误"))
			c.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(consts.PermissionAdmin) {
			utils.RespErr(ctx, c, ierror.NewIError(-1, "权限错误"))
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}

func GetClaimsFromCtx(ctx context.Context, c *app.RequestContext) (*utils.CustomClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, ierror.NewIError(-1, "未登录")
	}
	return claims.(*utils.CustomClaims), nil
}
