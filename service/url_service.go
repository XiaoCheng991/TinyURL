package service

import (
	"TinyURL/entity"
	"TinyURL/logic"
	"net/http"
)

// urlService 实现了 entity.URLService 接口
type urlService struct {
	repo entity.URLRepository
}

// NewURLService 构造函数，注入 repo 依赖
func NewURLService(repo entity.URLRepository) entity.URLService {
	return &urlService{repo: repo}
}

// Create 接口实现方法：创建
func (s *urlService) Create(longURL string) (*entity.URLMapping, int) {
	mapping, err := s.repo.Save(longURL)

	// 有错误 返回 500
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	// 否则返回 201
	return mapping, http.StatusCreated
}

// RedirectTo 接口实现方法：重定向
func (s *urlService) RedirectTo(shortCode string) (string, int) {
	id, err := logic.Decode(shortCode)
	if err != nil {
		return "", http.StatusInternalServerError
	}
	mapping, err := s.repo.FindByID(id)
	if err != nil {
		return "", http.StatusNotFound
	}
	return mapping.LongURL, http.StatusFound
}
