package entity

import "time"

// URLMapping 长短链接映射核心实体结构
// 贫血模型，仅包含数据库字段、数据库映射标签
type URLMapping struct {
	// ID 是数据库自增主键， 使用 uint64 支持海量数据存储
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// LongURL 是原来的长链，1，采用 text 类型
	LongURL string `gorm:"type:text;not null" json:"long_url"`

	// CreatedAt 是创建时间, 在 GROM 插入时自动填充
	CreatedAt time.Time `json:"created_at"`

	// ExpiredAt 记录过期时间（可选），企业级的链接时效管理
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// TableName 指向该实体对应的数据库表名，符合工程化命名规范
func (URLMapping) TableName() string { return "short_url_mappings" }

// ShortenRequest 创建短链接时的 API 请求载体
type ShortenRequest struct {
	LongURL    string `json:"long_url" binding:"required,url"`
	CustomCode string `json:"custom_code,omitempty"`
}

// ShortenResponse 返回给前端的统一数据格式
type ShortenResponse struct {
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
}
