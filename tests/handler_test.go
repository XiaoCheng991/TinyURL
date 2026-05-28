package tests

import (
	"TinyURL/gateway"
	"TinyURL/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"TinyURL/entity"
	"TinyURL/repo"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO 待改为Mock测试

// 每个测试前设置gin为测试模式，避免Debug日志刷屏
func init() {
	gin.SetMode(gin.TestMode)
}

// TestCreateShortURL_Success 测试创建短链
func TestCreateShortURL_Success(t *testing.T) {
	// 1. 准备：创建 repo 实例，构造handler
	r := repo.NewMemoryRepo()
	s := service.NewURLService(r)
	handler := gateway.CreateShortURL(s)

	// 2. 构造请求：body + NewRequest
	body := `{"long_url": "https://www.baidu.com"}`
	req := httptest.NewRequest("POST", "/api/v1/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// 3. 执行：怎么让 handler 跑起来
	router := gin.New()
	router.POST("/api/v1/shorten", handler)
	router.ServeHTTP(w, req)

	// 4. 断言：w.Code 应该是 201
	assert.Equal(t, http.StatusCreated, w.Code)

	// 解析响应体
	var resp entity.ShortenResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// 验证返回数据
	assert.Equal(t, "https://www.baidu.com", resp.LongURL)
	assert.NotEmpty(t, resp.ShortCode)
	assert.Contains(t, resp.ShortURL, resp.ShortCode)
	assert.False(t, resp.CreatedAt.IsZero())
}

func TestCreateShortURL_InvalidJSON(t *testing.T) {
	r := repo.NewMemoryRepo()
	s := service.NewURLService(r)
	handler := gateway.CreateShortURL(s)

	body := `这不是 json`
	req := httptest.NewRequest("POST", "/api/v1/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/api/v1/shorten", handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateShortURL_InvalidURL(t *testing.T) {
	r := repo.NewMemoryRepo()
	s := service.NewURLService(r)
	handler := gateway.CreateShortURL(s)

	// long_url不是合法URL，binding:"url"会拦截
	body := `{"long_url": "not-a-url"}`
	req := httptest.NewRequest("POST", "/api/v1/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/api/v1/shorten", handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRedirectURL_Success(t *testing.T) {
	r := repo.NewMemoryRepo()
	s := service.NewURLService(r)
	createHandler := gateway.CreateShortURL(s)
	redirectHandler := gateway.RedirectURL(s)

	// 先创建一条短链
	body := `{"long_url": "https://www.example.com"}`
	req := httptest.NewRequest("POST", "/api/v1/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	router.POST("/api/v1/shorten", createHandler)
	router.GET("/api/v1/:code", redirectHandler)
	router.ServeHTTP(w, req)

	// 解析拿到short_code
	var resp entity.ShortenResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}

	// 用short_code访问跳转
	req2 := httptest.NewRequest("GET", "/api/v1/"+resp.ShortCode, nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// 验证302重定向
	assert.Equal(t, http.StatusFound, w2.Code)
	assert.Equal(t, "https://www.example.com", w2.Header().Get("Location"))
}

func TestRedirectURL_NotFound(t *testing.T) {
	r := repo.NewMemoryRepo()
	s := service.NewURLService(r)
	handler := gateway.RedirectURL(s)

	// 访问一个不存在的短码
	req := httptest.NewRequest("GET", "/nonexist", nil)
	w := httptest.NewRecorder()

	router := gin.New()
	router.GET("/api/v1/:code", handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
