package ctl

import (
	"github.com/291700351/core"
	"github.com/gin-gonic/gin"
)

func A(engine *gin.Engine) {
	t := new(test)
	engine.GET("test", t.hello)
}

type test struct {
	*core.BaseController
}

func (t *test) hello(c *gin.Context) {
	t.Ok(c, false)
}
