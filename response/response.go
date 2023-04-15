package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response
// @Description 返回前端响应
// @Author John 2023-04-15 11:06:42
// @Param ctx
// @Param httpStatus HTTP状态码
// @Param code 前端自定义状态码
// @Param data 数据
// @Param msg 信息
func Response(ctx *gin.Context, httpStatus int, status int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"status": status,
		"data":   data,
		"msg":    msg,
	})
}

// Success
// @Description 请求成功响应
// @Author John 2023-04-15 11:17:00
// @Param ctx
// @Param data
// @Param msg
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// Fail
// @Description 请求错误响应
// @Author John 2023-04-15 11:17:11
// @Param ctx
// @Param data
// @Param msg
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}
