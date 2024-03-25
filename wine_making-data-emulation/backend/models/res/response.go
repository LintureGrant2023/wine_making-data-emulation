package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	Err     = 1
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 封装底层的响应
func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, "成功", data, c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, "成功", map[string]any{}, c)
}

func Error(msg string, c *gin.Context) {
	Result(Err, msg, map[string]any{}, c)
}
