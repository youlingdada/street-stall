package mdw

import (
	"github.com/youlingdada/street-stall/utils/redis_utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/youlingdada/street-stall/utils"
)

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录

func JwtMiddleware(accessUrls map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, ok := accessUrls[path]; ok {
			c.Next()
			return
		}

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "未登录", Data: nil}))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "auth格式错误,登录失效", Data: nil}))
			c.Abort()
			return
		}
		uuid, err := utils.ParseToken(authHeader)
		if err != nil {
			log.Printf("token解析失败，err->%s", err)
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "登录失效，请重新登录", Data: nil}))
			c.Abort()
			return
		}
		// 设置uuid,后面可以取到
		c.Set("uuid", *uuid)
		ttl, err := redis_utils.UserCache.Ttl(*uuid)
		if err != nil {
			log.Printf("获取用户登录过期时间失败，err->%s", err)
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "获取用户信息失败，请重新登录", Data: nil}))
			c.Abort()
			return
		}
		if ttl < 0 {
			log.Printf("登录过期，请重新登录,ttl->%d", ttl)
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "登录过期，请重新登录", Data: nil}))
			c.Abort()
			return
		}
		user, err := redis_utils.UserCache.Get(*uuid)
		if err != nil {
			log.Printf("获取用户信息失败，err->%s", err)
			c.JSON(http.StatusOK, utils.H(utils.Result{Code: 403, Message: "获取用户信息失败，请重新登录", Data: nil}))
			c.Abort()
			return
		}
		// 将用户信息存储在request 域中
		c.Set("loginUser", *user)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
