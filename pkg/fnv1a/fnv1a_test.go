package fnv1a

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestHash32(t *testing.T) {
	res, _ := hex.DecodeString("4f9f2cab")

	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), res},
	}

	hasher := NewFNV1a32()

	for _, tt := range tests {
		result := hasher.Hash(tt.input)

		if !bytes.Equal(result, tt.expected) {
			t.Errorf("Hash(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func BenchmarkHash64(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a64()

	for b.Loop() {
		hasher.Hash(data)
	}
}
