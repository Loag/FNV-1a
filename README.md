# FNV-1a Hash Implementation in Go

This package provides a simple and efficient implementation of the FNV-1a (Fowler-Noll-Vo) hash algorithm in Go. The implementation includes both 64-bit hash functions for byte slices and strings.

## Installation

```bash
go get github.com/yourusername/fnv1a
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/fnv1a"
)

func main() {
    // Hash a byte slice
    data := []byte("Hello, World!")
    hash1 := fnv1a.Hash64(data)
    fmt.Printf("Hash of %q: %d\n", data, hash1)

    // Hash a string
    text := "Hello, World!"
    hash2 := fnv1a.HashString64(text)
    fmt.Printf("Hash of %q: %d\n", text, hash2)
}
```

## Features

- 64-bit FNV-1a hash implementation
- Support for both byte slices and strings
- Thoroughly tested with benchmarks
- Zero dependencies

## Performance

The implementation is optimized for performance while maintaining simplicity. Benchmark results can be obtained by running:

```bash
go test -bench=.
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
