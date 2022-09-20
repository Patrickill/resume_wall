package main

import (
	"back/config"
	"back/db"
	"back/db/model"
	"back/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig() //初始化配置文件
	db.InitDB()         // 初始化数据库
	model.InitModel()   //初始化图表格
	if !config.Config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	router.InitRouter()                                //初始化路由
	router.Router.Run(":" + config.Config.Server.Port) //用gin框架启动服务器 并从配置文件读取开放端口
}
