package service

import (
	"TinyURL/entity"
	"TinyURL/logic"
	"TinyURL/repo"
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
func (s *urlService) Create(longURL string) (*entity.URLMapping, error) {
	mapping, err := s.repo.Save(longURL)

	// 有错误 返回 500
	if err != nil {
		return nil, repo.ErrSaveFailed
	}

	// 否则返回 201
	return mapping, nil
}

// RedirectTo 接口实现方法：重定向
func (s *urlService) RedirectTo(shortCode string) (string, error) {
	id, err := logic.Decode(shortCode)
	if err != nil {
		return "", logic.ErrDecodeFailed
	}
	mapping, err := s.repo.FindByID(id)
	if err != nil {
		return "", repo.ErrDataNotFound
	}
	return mapping.LongURL, nil
}

// GetInfo 根据短码查询映射详情
func (s *urlService) GetInfo(shortCode string) (*entity.URLMapping, error) {
	id, err := logic.Decode(shortCode)
	if err != nil {
		return nil, logic.ErrDecodeFailed
	}
	mapping, err := s.repo.FindByID(id)
	if err != nil {
		return nil, repo.ErrDataNotFound
	}
	return mapping, nil
}
