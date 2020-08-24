package v1

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/zhouwy1994/gin-example/pkg/ecode"
	"github.com/zhouwy1994/gin-example/pkg/gredis"
	"github.com/zhouwy1994/gin-example/pkg/logger"
	"github.com/zhouwy1994/gin-example/pkg/setting"

	"github.com/zhouwy1994/gin-example/models"
	"github.com/zhouwy1994/gin-example/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func UserRegister(c *gin.Context) {
	// 获取POST参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	code := ecode.SUCCESS
	data := make(map[string]interface{})

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	// 检查POST参数合法性
	if ok, _ := valid.Valid(&a); ok {
		// 检查用户是否已经存在
		if isExist := models.ExistUserByName(username); !isExist {
			// 用户信息写入数据库
			if err := models.AddUser(username, password); err == nil {
				// 生成Token
				if token, err := util.GenerateToken(username, password); err == nil {
					redisKey := fmt.Sprintf("%s_token", username)
					// Token写入Redis
					if err = gredis.Set(redisKey, token, setting.JwtExpireTime); err == nil {
						// 操作成功
						data["token"] = token
					} else {
						code = ecode.ERROR_AUTH_TOKEN
						logger.Error("USER", "token write to redis fail:%s", err)
					}
				} else {
					code = ecode.ERROR_AUTH_TOKEN
					logger.Error("USER", "generate token fail:%s", err)
				}
			} else {
				code = ecode.ERROR_DATABASE
				logger.Error("USER", "database write fail:%s", err)
			}
		} else {
			code = ecode.ERROR_USER_ALREADY_EXIST
		}
	} else {
		code = ecode.INVALID_PARAMS
		for _, err := range valid.Errors {
			logger.Warn("USER", "param parse error:%s:%s", err.Key, err.Message)
		}
	}

	(&app.Gin{c}).Response(200, code, data)
}

func UserLogin(c *gin.Context) {
	// 获取POST参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	code := ecode.SUCCESS
	data := make(map[string]interface{})

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	// 检查POST参数合法性
	if ok, _ := valid.Valid(&a); ok {
		// 检查用户是否已经存在
		if isValid,err := models.CheckUser(username, password); err == nil {
			// 用户验证通过
			if isValid {
				// 生成Token
				if token, err := util.GenerateToken(username, password); err == nil {
					redisKey := fmt.Sprintf("%s_token", username)
					// Token写入Redis
					if err = gredis.Set(redisKey, token, setting.JwtExpireTime); err == nil {
						// 操作成功
						data["token"] = token
					} else {
						code = ecode.ERROR_AUTH_TOKEN
						logger.Error("USER", "token write to redis fail:%s", err)
					}
				} else {
					code = ecode.ERROR_AUTH_TOKEN
					logger.Error("USER", "generate token fail:%s", err)
				}
			} else {
				code = ecode.ERROR_USER_AUTH
			}
		} else {
			code = ecode.ERROR_DATABASE
			logger.Error("USER", "database fail:%s", err)
		}
	} else {
		code = ecode.INVALID_PARAMS
		for _, err := range valid.Errors {
			logger.Warn("USER", "param parse error:%s:%s", err.Key, err.Message)
		}
	}

	(&app.Gin{c}).Response(200, code, data)
}
