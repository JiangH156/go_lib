package controller

import (
	"Go_lib/response"
	"Go_lib/service"
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

func NewReportController() ReportController {
	return ReportController{}
}
