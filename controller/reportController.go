package controller

import (
	"Go_lib/model"
	"Go_lib/response"
	"Go_lib/service"
	"Go_lib/utils"
	"Go_lib/vo"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
}

// GetReportRecords
// @Description  获取举报记录
// @Author John 2023-04-26 19:13:25
// @Param ctx
func (r *ReportController) GetReportRecords(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")

	reportService := service.NewReportService()
	reports, lErr := reportService.GetReportRecords(readerId)

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
		"msg":    "获取举报记录成功",
		"data":   reports,
	})
}

// GetAllReportRecords
// @Description 管理员获取所有举报记录
// @Author John 2023-04-28 15:09:49
// @Param ctx
func (r *ReportController) GetAllReportRecords(ctx *gin.Context) {
	reportService := service.NewReportService()
	reportVos, lErr := reportService.GetAllReportRecords()

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
		"msg":    "获取举报记录成功",
		"data":   reportVos,
	})
}

// CreateReport
// @Description 用户举报评论
// @Author John 2023-05-02 11:41:32
// @Param ctx
func (r *ReportController) CreateReport(ctx *gin.Context) {
	// 数据接收
	commentId := ctx.PostForm("commentId")
	reporterId := ctx.PostForm("reporterId")

	reportService := service.NewReportService()
	lErr := reportService.CreateReport(commentId, reporterId)

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
		"msg":    "获取举报记录成功",
	})
}

// ManageReport
// @Description 管理员处理举报评论功能、删除评论功能入口
// @Author John 2023-05-02 16:17:11
// @Param ctx
func (r *ReportController) ManageReport(ctx *gin.Context) {
	// 数据接收
	commentId := ctx.PostForm("commentId")
	reporterId := ctx.PostForm("reporterId")
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	reportDate := ctx.PostForm("reportdate")
	date := ctx.PostForm("date")
	status := ctx.PostForm("status")

	t, _ := utils.ParseTime(date)
	tt, _ := utils.ParseTime(reportDate)
	// 封装举报信息
	reportInfo := vo.ReportVo{
		CommentID:  commentId,
		ReporterID: reporterId,
		ReaderID:   readerId,
		ReportDate: model.Time(tt),
		BookID:     bookId,
		Date:       model.Time(t),
		Status:     status,
	}

	reportService := service.NewReportService()
	lErr := reportService.ManageReport(reportInfo)
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
		"msg":    "请求成功",
	})
}

func NewReportController() ReportController {
	return ReportController{}
}
