package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler(c *gin.Context) {
	defer func() {
		r := recover()
		if nil != r {
			c.JSON(http.StatusOK, JsonFail(500, "服务器异常"))
			c.Abort()
		}
	}()
	c.Next()

}

func NoRouter(c *gin.Context) {
	c.JSON(http.StatusOK, JsonFail(404, "资源未找到"))
	c.Abort()

}
