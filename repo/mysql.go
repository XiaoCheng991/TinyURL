package repo

import (
	"TinyURL/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 持久层
type mysqlRepo struct {
	db *gorm.DB
}

// DSN 格式：root:admin@tcp(127.0.0.1:3306)/tinyurl?charset=utf8mb4&parseTime=True&loc=Local

// NewMySQLRepo 构造方法
func NewMySQLRepo(dsn string) (entity.URLRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&entity.URLMapping{}); err != nil {
		return nil, err
	}
	return &mysqlRepo{db: db}, nil
}

// Save 保存
func (r *mysqlRepo) Save(longURL string) (*entity.URLMapping, error) {
	mapping := &entity.URLMapping{
		LongURL: longURL,
	}
	if err := r.db.Create(mapping).Error; err != nil {
		return nil, err
	}
	return mapping, nil
}

// FindByID 根据ID查询
func (r *mysqlRepo) FindByID(id uint64) (*entity.URLMapping, error) {
	var mapping entity.URLMapping
	if err := r.db.First(&mapping, id).Error; err != nil {
		return nil, err
	}
	return &mapping, nil
}
