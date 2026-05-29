package logic

import "errors"

var (
	// ErrInvalidBase62Char 非法 Base62 字符
	ErrInvalidBase62Char = errors.New("invalid base62 character")

	// ErrEmptyShortCode 空短码
	ErrEmptyShortCode = errors.New("empty short code")

	// ErrCodeTooLong 短码过长
	ErrCodeTooLong = errors.New("short code too long")

	// ErrURLNotFound URL 未找到
	ErrURLNotFound = errors.New("url not found")

	// ErrKeyNotFound Key 未找到
	ErrKeyNotFound = errors.New("key not found")
)
