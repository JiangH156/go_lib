package main

import (
	"fmt"
	"github.com/John/Go_lib/common"
	"github.com/John/Go_lib/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	r := gin.Default()
	InitConfig()
	common.InitDB()
	db := common.GetDB()
	defer func() {
		sqlDB, _ := db.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	// 路由收集
	r = router.CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run(":8080")
	}

}

// InitConfig
// @Description 初始化配置
// @Author John 2023-04-13 22:26:53 ${time}
func InitConfig() {
	dir, _ := os.Getwd()
	// 配置文件所在目录
	viper.AddConfigPath(dir + "/config")
	// 配置文件名（不带后缀）
	viper.SetConfigName("application")
	// 配置文件类型
	viper.SetConfigType("yml")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to load config file: %s\n", err))
	}
}
