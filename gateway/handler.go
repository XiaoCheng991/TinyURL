package gateway

import (
	"TinyURL/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateShortURL 新建短链
func CreateShortURL(c *gin.Context) {
	// 1. 绑定请求体 JSON 到 ShortenRequest
	var req entity.ShortenRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		// 绑定失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 验证参数（Gin binding 帮你做了一部分，但你还要检查空值）

	// 返回 JSON
	c.JSON(http.StatusOK, gin.H{
		"short_url":  "http://tiny.url/2bg",
		"short_code": "2bg",
		"long_url":   req.LongURL,
	})

	// 3. 调 logic/repo 处理业务（现在还没 repo 假装存在）

	// 4. 返回响应

}

// RedirectURL 重定向
func RedirectURL(c *gin.Context) {
	code := c.Param("code")
	_ = code // TODO: 查repo获取长链后重定向

	// HTTP 重定向
	c.Redirect(http.StatusFound, "https://www.baidu.com")
}
