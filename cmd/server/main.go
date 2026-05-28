package main

import (
	"TinyURL/gateway"
	"TinyURL/repo"
	"TinyURL/service"
	"fmt"
	"log"

	"TinyURL/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode) // 生产模式

	// 创建 Gin 路由
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// 基础健康检查
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 初始化MySQL repo
	mysqlRepo, err := repo.NewMySQLRepo(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("Failed to create repo: %v", err)
	}

	// 初始化 service（注入repo）
	urlService := service.NewURLService(mysqlRepo)

	// 注册handler（注入 service）
	r.POST("/api/v1/shorten", gateway.CreateShortURL(urlService))
	r.GET("/api/v1/:code", gateway.RedirectURL(urlService))

	// 启动服务
	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
