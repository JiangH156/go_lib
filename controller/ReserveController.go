package controller

import (
	"Go_lib/model"
	"Go_lib/response"
	"Go_lib/service"
	"Go_lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReserveController struct {
}

// CreateReserveRecord
// @Description 添加预约记录
// @Author John 2023-04-19 18:52:17
// @Param ctx
func (r *ReserveController) CreateReserveRecord(ctx *gin.Context) {
	var reserveService = service.NewReserveService()
	// 数据接收
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	date := ctx.PostForm("date")
	status := ctx.DefaultPostForm("status", "已预约")

	// 格式化时间
	addTime, err := utils.ParseTime(date)
	if err != nil {
		fmt.Println("时间格式化失败")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "预约失败",
		})
		return
	}

	var addReserve = model.Reserve{
		ReaderId: readerId,
		BookId:   bookId,
		Date:     addTime,
		Status:   status,
	}

	// 新增记录
	lErr := reserveService.CreateReserveRecord(addReserve)

	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}

	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "预约成功",
	})
}

// GetReserves
// @Description 获取预约信息
// @Author John 2023-04-19 18:53:26
// @Param ctx
func (r *ReserveController) GetReserves(ctx *gin.Context) {
	var reserveService = service.NewReserveService()
	readerId := ctx.PostForm("readerId")
	if readerId == "" {
		fmt.Println("readerId为空")
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    "查询预约记录失败",
		})
		return
	}

	reserveVOs, lErr := reserveService.GetReserves(readerId)
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}

	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "读者请求预约记录成功",
		"data":   reserveVOs,
	})

}

// DeleteReserveRecord
// @Description 取消预约记录接口
// @Author John 2023-04-19 23:04:01
// @Param ctx
func (r *ReserveController) DeleteReserveRecord(ctx *gin.Context) {
	var reserveService = service.NewReserveService()
	bookId := ctx.PostForm("bookId")
	readerId := ctx.PostForm("readerId")
	date := ctx.PostForm("date")

	ddate, err := utils.ParseTime(date)
	if err != nil {
		fmt.Println("时间格式化失败")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
		})
		return
	}
	delReserve := model.Reserve{
		BookId:   bookId,
		ReaderId: readerId,
		Date:     ddate,
	}
	lErr := reserveService.DeleteReserveRecord(delReserve)
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}

	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "取消预约成功!",
	})
}
func NewReserveController() ReserveController {
	return ReserveController{}
}
