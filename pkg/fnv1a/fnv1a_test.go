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

func TestHash64(t *testing.T) {
	res, _ := hex.DecodeString("a430d84680aabd0b")

	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), res},
	}

	hasher := NewFNV1a64()

	for _, tt := range tests {
		result := hasher.Hash(tt.input)

		if !bytes.Equal(result, tt.expected) {
			t.Errorf("Hash(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func BenchmarkHash32(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a32()

	for b.Loop() {
		hasher.Hash(data)
	}
}

func BenchmarkHash64(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a64()

	for b.Loop() {
		hasher.Hash(data)
	}
}

func BenchmarkHash128(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a128()

	for b.Loop() {
		hasher.Hash(data)
	}
}

func BenchmarkHash256(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a256()
	for b.Loop() {
		hasher.Hash(data)
	}
}

func BenchmarkHash512(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a512()
	for b.Loop() {
		hasher.Hash(data)
	}
}

func BenchmarkHash1024(b *testing.B) {
	data := []byte("Hello")
	hasher := NewFNV1a1024()
	for b.Loop() {
		hasher.Hash(data)
	}
}
