# FNV-1a

contains all implementations, 32 - 1024 using go big.Int

## Example 

``` go

hasher := NewFNV1a32()
result := hasher.Hash([]byte("hello"))

```

## Benchmarks

- ns/op 590
- allocations 7

## TODO

clean up allocations more. most likley by removing big int