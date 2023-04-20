package controller

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/response"
	"Go_lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ReserveController struct {
	DB *gorm.DB
}

// AddReserve
// @Description 添加预约记录
// @Author John 2023-04-19 18:52:17
// @Param ctx
func (r *ReserveController) AddReserve(ctx *gin.Context) {
	var reserveRepository = repository.NewReserveRepository()
	// 数据接收
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	date := ctx.PostForm("date")
	status := ctx.DefaultPostForm("status", "已预约")

	if readerId == "" || bookId == "" {
		fmt.Println("预约失败")
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    "预约失败",
		})
		return
	}

	// 验证数据库是否已经存在该预约
	var reserve = model.Reserve{}
	r.DB.Where("reader_id = ?", readerId).Where("book_id = ?", bookId).First(&reserve)
	if reserve.Id != "" {
		fmt.Println("预约记录已存在")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "预约失败",
		})
		return
	}

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
	//fmt.Println(addTime)

	addReserve := model.Reserve{
		ReaderId: readerId,
		BookId:   bookId,
		Date:     addTime,
		Status:   status,
	}
	if err := reserveRepository.AddReserve(addReserve); err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "预约失败",
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "预约成功",
	})
}

// QueryReserveByReaderId
// @Description 读者请求预约记录
// @Author John 2023-04-19 18:53:26
// @Param ctx
func (r *ReserveController) QueryReserveByReaderId(ctx *gin.Context) {
	var reserveRepository = repository.NewReserveRepository()
	readerId := ctx.PostForm("readerId")
	if readerId == "" {
		fmt.Println("readerId为空")
		response.Fail(ctx, gin.H{
			"status": 400,
			"msg":    "查询预约记录失败",
		})
		return
	}

	reserveVOs, err := reserveRepository.QueryReserveByReaderId(readerId)
	if err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "读者请求预约记录失败",
		})
		return
	}

	//查询数据为空
	if len(reserveVOs) == 0 {
		fmt.Println("读者请求预约记录为空")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "读者请求预约记录为空",
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "读者请求预约记录成功",
		"data":   reserveVOs,
	})

}

// DeleteReserve
// @Description 取消预约记录接口
// @Author John 2023-04-19 23:04:01
// @Param ctx
func (r *ReserveController) DeleteReserve(ctx *gin.Context) {
	var reserveRepository = repository.NewReserveRepository()
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

	if err := reserveRepository.DeleteReserve(bookId, readerId, ddate); err != nil {
		fmt.Println("取消预约失败")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "取消预约失败",
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "取消预约成功!",
	})
}
func NewReserveController() ReserveController {
	return ReserveController{
		DB: common.GetDB(),
	}
}
