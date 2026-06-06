package repo

import "errors"

var (
	// ErrInvalidURL 无效URL
	ErrInvalidURL = errors.New("invalid URL")

	// ErrInvalidCode 无效编码
	ErrInvalidCode = errors.New("invalid code")

	// ErrDataNotFound 数据未找到
	ErrDataNotFound = errors.New("data not found")

	// ErrSaveFailed 保存到数据库失败
	ErrSaveFailed = errors.New("database save failed")
)
