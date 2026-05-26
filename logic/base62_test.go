package logic

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		id   uint64
		want string
	}{
		{name: "id=0", id: 0, want: "0"},
		{name: "id=1", id: 1, want: "1"},
		{name: "id=9", id: 9, want: "9"},
		{name: "id=10", id: 10, want: "A"},
		{name: "id=61", id: 61, want: "z"},
		{name: "id=62", id: 62, want: "10"},
		{name: "id=10024", id: 10024, want: "2bg"},
		{name: "id=999999", id: 999999, want: "4C91"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.id)
			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		want      uint64
		shouldErr bool
	}{
		{name: "code=0", code: "0", want: 0, shouldErr: false},
		{name: "code=1", code: "1", want: 1, shouldErr: false},
		{name: "code=9", code: "9", want: 9, shouldErr: false},
		{name: "code=A", code: "A", want: 10, shouldErr: false},
		{name: "code=z", code: "z", want: 61, shouldErr: false},
		{name: "code=10", code: "10", want: 62, shouldErr: false},
		{name: "code=2bg", code: "2bg", want: 10024, shouldErr: false},
		{name: "code=4C91", code: "4C91", want: 999999, shouldErr: false},
		{name: "code=invalid code", code: "!!!", want: 0, shouldErr: true},
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
			if result != tt.want {
				t.Errorf("Decode() = %v, want %v", result, tt.want)
			}
		})
	}
}
