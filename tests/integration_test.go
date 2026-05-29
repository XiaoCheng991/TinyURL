package tests

import (
	"TinyURL/config"
	"TinyURL/gateway"
	"TinyURL/repo"
	"TinyURL/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_CreateAndRedirect 测试创建并重定向
func TestIntegration_CreateAndRedirect(t *testing.T) {
	// 1. 加载配置， 初始化 MySQL repo
	cfg, err := config.Load("../config/config.yaml")
	assert.NoError(t, err)

	mysqlRepo, err := repo.NewMySQLRepo(cfg.Database.DSN())
	assert.NoError(t, err)

	// 2. 初始化 service + handler
	uslService := service.NewURLService(mysqlRepo)
	createHandler := gateway.CreateShortURL(uslService)
	redirectHandler := gateway.RedirectURL(uslService)

	// 3. 起真实 server
	router := gin.New()
	router.POST("/api/v1/shorten", createHandler)
	router.GET("/api/v1/:code", redirectHandler)
	server := httptest.NewServer(router)
	defer server.Close()

	// 4. 发 create 请求
	body := `{"long_url": "https://www.baidu.com"}`
	resp, err := http.Post(server.URL+"/api/v1/shorten", "/application/json", strings.NewReader(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createResp struct {
		LongURL   string `json:"long_url"`
		ShortCode string `json:"short_code"`
	}
	err = json.NewDecoder(resp.Body).Decode(&createResp)
	if err != nil {
		return
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}

	assert.Equal(t, "https://www.baidu.com", createResp.LongURL)
	assert.NotEmpty(t, createResp.ShortCode)

	// 5. 发 redirect 请求
	redirectURL := fmt.Sprintf("%s/api/v1/%s", server.URL, createResp.ShortCode)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // 不跟随重定向
		},
	}
	resp2, err := client.Get(redirectURL)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp2.Body)

	assert.Equal(t, http.StatusFound, resp2.StatusCode)
	assert.Equal(t, "https://www.baidu.com", resp2.Header.Get("Location"))

	t.Logf("create shortCode=%s, redirect Location=%s",
		createResp.ShortCode, resp2.Header.Get("Location"))
}
