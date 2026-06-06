package entity

// URLRepository URL仓库层接口
type URLRepository interface {
	Save(longURL string) (*URLMapping, error)
	FindByID(id uint64) (*URLMapping, error)
}

// URLService URL服务层接口
type URLService interface {
	Create(longURL string) (*URLMapping, error)
	RedirectTo(shortCode string) (string, error)
	GetInfo(shortCode string) (*URLMapping, error)
}
