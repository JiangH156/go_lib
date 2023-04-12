package main

import (
	"WeChat/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TableRecord struct {
	Name     string `json:"name"`
	TodayBuy int    `json:"todayBuy"`
	MonthBuy int    `json:"monthBuy"`
	TotalBuy int    `json:"totalBuy"`
}

type VideoData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	r := gin.Default()
	cors := middleware.CORSMiddleware()
	r.Use(cors)
	data := struct {
		TableData []TableRecord `json:"tableData"`
		VideoData []VideoData   `json:"videoData"`
	}{
		[]TableRecord{
			{"oppo", 500, 3500, 22000},
			{"vivo", 300, 2200, 24000},
			{"苹果", 800, 4500, 65000},
			{"小米", 1200, 6500, 45000},
			{"三星", 300, 2000, 34000},
			{"魅族", 350, 3000, 22000},
		},
		[]VideoData{
			{
				"小米",
				2999,
			},
			{
				"苹果",
				5999,
			},
			{
				"vivo",
				1500,
			},
			{
				"oppo",
				1999,
			},
			{
				"魅族",
				2200,
			},
			{
				"三星",
				4500,
			},
		},
	}

	r.GET("/api/home/getData", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "20000",
			"data": data,
		})

	})
	r.Run(":8081")
}
