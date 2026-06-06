package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Recovery 替代 gin.Recovery()，带错误栈输出
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[PANIC] %s %s: %v\n", c.Request.Method, c.Request.URL.Path, err)
				c.AbortWithStatusJSON(500, gin.H{
					"error": "internal server error",
				})
			}
		}()
		c.Next()
	}
}

// RequestLogger 自定义请求日志中间件
// 格式: [2026-06-04 15:30:00] 200 POST /api/v1/shorten 12.3ms
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf("[%s] %d %s %s %.1fms\n",
			start.Format("2006-01-02 15:04:05"),
			status,
			method,
			path,
			float64(latency.Microseconds())/1000.0,
		)
	}
}

// ErrorHandler 统一错误处理中间件
// 捕获 handler 中通过 c.Error() 设置的错误
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			fmt.Printf("[ERROR] %s: %v\n", c.Request.URL.Path, err.Err)
		}
	}
}
