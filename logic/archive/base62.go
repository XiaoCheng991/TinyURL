package archive

import "TinyURL/logic"

// Base62 字符集：0-9, a-z, A-Z
// 共 62 个字符，索引 0-61
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Base62CharsLength Base62 字符集长度
const Base62CharsLength = len(base62Chars)

// Encode 将 uint64 ID 转换为 Base62 短码
// 例如: Encode(10024) -> "2Bk"
func Encode(id uint64) string {
	if id == 0 {
		return string(base62Chars[0])
	}

	// 用于存储余数（逆序）
	var chars []byte

	for id > 0 {
		// 取余得到当前位的字符
		// 注意：需要转换为 int 作为索引
		remainder := int(id % uint64(Base62CharsLength))
		// 转换为对应字符并追加
		chars = append(chars, base62Chars[remainder])
		// 除以 62，继续处理下一位
		id /= uint64(Base62CharsLength)
	}

	// 反转数组（因为我们是从低位开始取的）
	// Java: Collections.reverse()
	// Go: 交换首尾元素
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// Decode 将 Base62 短码还原为 uint64 ID
// 例如: Decode("2Bk") -> 10024
func Decode(s string) (uint64, error) {
	var result uint64

	for _, c := range s {
		// 查找字符在字符集中的位置
		idx := -1
		for i, ch := range base62Chars {
			if ch == c {
				idx = i
				break
			}
		}

		// 未找到对应字符
		if idx == -1 {
			return 0, logic.ErrInvalidBase62Char
		}

		// result = result * 62 + idx
		// 对应十进制的 123 = ((1 * 10) + 2) * 10 + 3
		result = result*uint64(Base62CharsLength) + uint64(idx)
	}

	return result, nil
}

// MustDecode 类似 Decode，但失败时 panic（类似 strconv.Atoi）
func MustDecode(s string) uint64 {
	id, err := Decode(s)
	if err != nil {
		panic(err)
	}
	return id
}
