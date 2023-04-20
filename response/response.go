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
func Response(ctx *gin.Context, httpStatus int, data gin.H) {
	ctx.JSON(httpStatus, data)
}

// Success
// @Description 请求成功响应
// @Author John 2023-04-15 11:17:00
// @Param ctx
// @Param data
// @Param msg
func Success(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusOK, data)
}

// Fail
// @Description 请求错误响应
// @Author John 2023-04-15 11:17:11
// @Param ctx
// @Param data
// @Param msg
func Fail(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusBadRequest, data)
}
