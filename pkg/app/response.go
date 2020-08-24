package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouwy1994/gin-example/pkg/ecode"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  ecode.GetMsg(errCode),
		"data": data,
	})
}