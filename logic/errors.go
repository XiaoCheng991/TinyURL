package logic

import "errors"

var (
	// ErrInvalidBase62Char 非法 Base62 字符
	ErrInvalidBase62Char = errors.New("invalid base62 character")

	// ErrEmptyShortCode 空短码
	ErrEmptyShortCode = errors.New("empty short code")

	// ErrCodeTooLong 短码过长
	ErrCodeTooLong = errors.New("short code too long")

	// ErrDecodeFailed 解码失败
	ErrDecodeFailed = errors.New("decode failed")
)
