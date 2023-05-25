package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (bc *BaseController) Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, JsonSuccess(data))
}
func (bc *BaseController) Fail(c *gin.Context, code int, msg string) {
	bc.FailWithCode(c, http.StatusOK, code, msg)
}

func (bc *BaseController) FailWithCode(c *gin.Context, httpCode int, code int, msg string) {
	c.JSON(httpCode, JsonFail(code, msg))
}
