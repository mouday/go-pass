package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/handler"
	"github.com/mouday/cron-admin/src/router"
	"github.com/mouday/cron-admin/src/service"
)

//go:embed public/*
var local embed.FS

func main() {
	// app
	env := config.GetEnv()
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.SetTrustedProxies(nil)

	// 全局异常捕获
	app.Use(handler.Recover)

	app.Use(handler.AuthMiddleware())

	// 注册路由
	router.RegistRouter(app)

	// 数据库迁移
	config.Migrate()

	// 初始化数据
	config.InitData()

	// 初始化定时任务
	service.InitCron()

	// 启动消费者
	go service.Consumer()

	// 【Go语言】gin + go:embed 打包静态资源文件
	// ref: https://blog.csdn.net/Regulations/article/details/128858670
	fp, _ := fs.Sub(local, "public")
	app.StaticFS("/", http.FS(fp))

	fmt.Println("**********************")
	fmt.Println("Welcome Use Cron Admin")
	fmt.Println("server run at: ", "http://127.0.0.1:8082")
	fmt.Println("**********************")

	// 监听并在 http://127.0.0.1:8082 上启动服务
	err := app.Run("127.0.0.1:8082")

	if err != nil {
		fmt.Println(err)
	}
}
