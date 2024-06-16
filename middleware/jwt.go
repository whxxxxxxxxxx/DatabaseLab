package middleware

import (
	"DatabaseLab/middleware/response"
	"DatabaseLab/pkg/ctl"
	"DatabaseLab/pkg/errorx"
	"DatabaseLab/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = errorx.Success
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			code = errorx.InvalidParams
			c.JSON(200, gin.H{
				"status": code,
				"msg":    errorx.GetMsg(code),
				"data":   "Token不能为空",
			})
			c.Abort()
			return
		}
		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = errorx.ErrorAuthCheckTokenFail
		}
		if code != errorx.Success {
			res := response.JsonResponse{
				Code:    code,
				Message: errorx.GetMsg(code),
				Data:    "鉴权失败",
				Error:   err,
			}
			response.HttpResponse(c, res)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(newAccessToken)
		if err != nil {
			code = errorx.ErrorAuthCheckTokenFail
			res := response.JsonResponse{
				Code:    code,
				Message: errorx.GetMsg(code),
				Data:    err.Error(),
			}
			response.HttpResponse(c, res)
			c.Abort()
			return
		}
		SetToken(c, newAccessToken, newRefreshToken)
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	c.Header("access_token", accessToken)
	c.Header("refresh_token", refreshToken)
	c.SetCookie("access_token", accessToken, 3600*24, "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24, "/", "", true, true)
}
