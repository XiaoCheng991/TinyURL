package entity

import (
	"time"
)

// URLMapping 定义了长短链接映射的核心实体结构
// 该结构体采用了贫血模型设计，仅包含数据字段和数据库映射标签 [1]
type URLMapping struct {
	// ID 是数据库自增主键，建议使用 uint64 以支持海量数据存储 [1]
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// ShortCode 是生成的短码（如 "6pWz"），必须建立唯一索引以保证跳转唯一性 [1]
	ShortCode string `gorm:"uniqueIndex;type:varchar(10);not null" json:"short_code"`

	// LongURL 存储原始长链接，由于 URL 可能很长，建议使用 text 类型 [1]
	LongURL string `gorm:"type:text;not null" json:"long_url"`

	// CreatedAt 记录创建时间，由 GORM 在插入时自动填充 [1]
	CreatedAt time.Time `json:"created_at"`

	// ExpiredAt 记录过期时间（可选），用于企业级的链接时效管理
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// TableName 指定该实体对应的数据库表名，符合工程化命名规范
func (URLMapping) TableName() string {
	return "short_url_mappings"
}

// ---------------------------------------------------------
// 补充：为了提升代码可读性，通常在 entity 层也会定义相关的请求/响应结构
// ---------------------------------------------------------

// ShortenRequest 定义了创建短链接时的 API 请求载体
type ShortenRequest struct {
	LongURL    string `json:"long_url" binding:"required,url"` // 使用 gin 验证器校验 URL 合法性
	CustomCode string `json:"custom_code,omitempty"`           // 用户自定义短码（可选）
}

// ShortenResponse 定义了返回给前端的统一数据格式
type ShortenResponse struct {
	ShortURL  string    `json:"short_url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
}
