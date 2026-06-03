package logic

// base62 字符集常量 【0-9 A-Z a-z】
// 62个字符，索引 0-61
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Base62Length base62 字符长度常量
const Base62Length = 62

// ASCII 码对应的数组
var base62Lookup [128]int8

func init() {
	// 先把所有位置设为 -1（非法标记）
	for i := range base62Lookup {
		base62Lookup[i] = -1
	}
	// 只填写 base62 中的 62 个字符
	for i := 0; i < len(base62Chars); i++ {
		// base62Chars[i] 得到的是 base62 对应的ASCII 码值，比如k是107
		base62Lookup[base62Chars[i]] = int8(i)
	}
}

// Encode 将 uint64 ID 转换为 Base62 短码
// 例如: Encode(10024) -> "2bg"
func Encode(id uint64) string {
	if id == 0 {
		return string(base62Chars[0])
	}

	// 用于存储余数（逆序）
	chars := make([]byte, 0, 12)

	// 拿到数据库对应行的索引 开始转换
	for id > 0 {
		// 循环的去除以 62 的到 商 和 余数
		remainder := int(id % uint64(Base62Length))

		// 转换为对应字符并追加
		chars = append(chars, base62Chars[remainder])

		// 除以字符长度，继续下一位
		id /= uint64(Base62Length)
	}

	// 现在 chars 收集的结果是倒序的，需要反转字符串
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// Decode 解码, 将 Base62 短码 还原为 uint64 的 Id
// 例如: Decode("2bg") -> 10024
func Decode(s string) (uint64, error) {
	var res uint64

	// c 是 rune 类型的 可以存储一个完整的 unicode 字符
	for i := 0; i < len(s); i++ {
		// 当前值
		c := s[i]
		if c >= 128 {
			return 0, ErrInvalidBase62Char
		}

		// O(1) 查表，拿到索引
		idx := base62Lookup[c]
		if idx == -1 {
			return 0, ErrInvalidBase62Char
		}

		// 62进制 res = res * 62 + idx
		// 十进制的 123 = ((1 * 10) + 2) * 10 + 3
		res = res*uint64(Base62Length) + uint64(idx)
	}
	return res, nil
}

// MustDecode 类似 Decode，但失败时 panic（类似 strconv.Atoi）
func MustDecode(s string) uint64 {
	id, err := Decode(s)
	if err != nil {
		panic(err)
	}
	return id
}
