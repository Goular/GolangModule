package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ApiServer/demo06/pkg/errno"
)

// 统一响应的结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 发送统一的响应
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
