package gateway

import (
	"TinyURL/entity"
	"TinyURL/logic"
	"TinyURL/repo"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateShortURL 返回一个闭包，注入repo依赖
func CreateShortURL(s entity.URLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 绑定请求体 JSON 到 ShortenRequest
		var req entity.ShortenRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			// 绑定失败，返回错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		// 2. 调 logic/repo 保存，拿到自增ID
		mapping, err := s.Create(req.LongURL)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, repo.ErrSaveFailed) {
				status = http.StatusInternalServerError
			}
			c.JSON(status, gin.H{"error": "create failed"})
			return
		}

		shortCode := logic.Encode(mapping.ID)

		// 3. 返回
		c.JSON(http.StatusCreated, entity.ShortenResponse{
			LongURL:   mapping.LongURL,
			ShortCode: shortCode,
			ShortURL:  "http://tiny.url/" + shortCode,
			CreatedAt: mapping.CreatedAt,
		})
	}
}

// RedirectURL 重定向
func RedirectURL(s entity.URLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
		code := c.Param("code")

		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing short_code"})
			return
		}

		// 查 repo 拿长链
		longURL, err := s.RedirectTo(code)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, logic.ErrDecodeFailed) {
				status = http.StatusInternalServerError
			}
			if errors.Is(err, repo.ErrDataNotFound) {
				status = http.StatusInternalServerError
			}
			c.JSON(status, gin.H{"error": "not found"})
			return
		}

		c.Redirect(http.StatusFound, longURL)
	}
}

// GetURLInfo 查询短链详情
func GetURLInfo(s entity.URLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing short_code"})
			return
		}

		mapping, err := s.GetInfo(code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		c.JSON(http.StatusOK, entity.ShortenResponse{
			LongURL:   mapping.LongURL,
			ShortCode: code,
			ShortURL:  "http://tiny.url/" + code,
			CreatedAt: mapping.CreatedAt,
		})
	}
}
