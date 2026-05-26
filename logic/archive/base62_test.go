package archive

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		id       uint64
		expected string
	}{
		{"zero", 0, "0"},
		{"one", 1, "1"},
		{"nine", 9, "9"},
		{"ten", 10, "A"},
		{"small", 61, "z"},
		{"62", 62, "10"},
		{"example", 10024, "2bg"},
		{"large", 999999, "4C91"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Encode(tt.id)
			if result != tt.expected {
				t.Errorf("Encode(%d) = %q, want %q", tt.id, result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		expected  uint64
		shouldErr bool
	}{
		{"zero", "0", 0, false},
		{"one", "1", 1, false},
		{"nine", "9", 9, false},
		{"ten", "A", 10, false},
		{"small", "z", 61, false},
		{"62", "10", 62, false},
		{"example", "2bg", 10024, false},
		{"large", "4C91", 999999, false},
		{"invalid char", "!!!", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.code)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("Decode(%q) should return error", tt.code)
				}
				return
			}
			if err != nil {
				t.Errorf("Decode(%q) unexpected error: %v", tt.code, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Decode(%q) = %d, want %d", tt.code, result, tt.expected)
			}
		})
	}
}

// TestEncodeDecodeRoundTrip 测试编解码往返（重要！）
func TestEncodeDecodeRoundTrip(t *testing.T) {
	ids := []uint64{0, 1, 10, 61, 62, 100, 1000, 10024, 100000, 999999, 1000000}

	for _, id := range ids {
		code := Encode(id)
		decoded, err := Decode(code)
		if err != nil {
			t.Errorf("Decode(%s) failed: %v", code, err)
			continue
		}
		if decoded != id {
			t.Errorf("Round trip failed: Encode(%d) = %s, Decode = %d", id, code, decoded)
		}
	}
}

func TestMustDecode(t *testing.T) {
	result := MustDecode("2bg")
	if result != 10024 {
		t.Errorf("MustDecode(\"2bg\") = %d, want 10024", result)
	}

	// 应该 panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustDecode with invalid input should panic")
		}
	}()
	MustDecode("!!!")
}

// BenchmarkEncode 基准测试
func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(1000000)
	}
}

// BenchmarkDecode 基准测试
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("4C91")
	}
}
