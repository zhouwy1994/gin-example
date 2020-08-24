package jwt

import (
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/zhouwy1994/gin-example/pkg/ecode"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = ecode.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = ecode.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = ecode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = ecode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
