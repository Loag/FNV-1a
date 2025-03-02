package fnv1a

import (
	"encoding/hex"
	"math/big"
)

var (
	offset32   = "811c9dc5"
	prime32    = "01000193"
	offset64   = "cbf29ce484222325"
	prime64    = "00000100000001b3"
	offset128  = "6c62272e07bb014262b821756295c58d"
	prime128   = "0000000001000000000000000000013b"
	offset256  = "dd268dbcaac550362d98c384c4e576ccc8b1536847b6bbb31023b4c8caee0535"
	prime256   = "0000000000000000000001000000000000000000000000000000000000000163"
	offset512  = "b86db0b1171f4416dca1e50f309990acac87d059c90000000000000000000d21e948f68a34c192f62ea79bc942dbe7ce182036415f56e34bac982aac4afe9fd9"
	prime512   = "00000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000157"
	offset1024 = "0000000000000000005f7a76758ecc4d32e56d5a591028b74b29fc4223fdada16c3bf34eda3674da9a21d9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004c6d7eb6e73802734510a555f256cc005ae556bde8cc9c6a93b21aff4b16c71ee90b3"
	prime1024  = "000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018d"
)

func hexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

type Hasher interface {
	Hash(data []byte) []byte
}

type FNV1a struct {
	offset *big.Int
	prime  *big.Int
	mod    *big.Int
}

func NewFNV1a32() Hasher {
	offset, err := hexToBytes(offset32)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime32)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 32)
}

func NewFNV1a64() Hasher {
	offset, err := hexToBytes(offset64)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime64)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 64)
}

func NewFNV1a128() Hasher {
	offset, err := hexToBytes(offset128)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime128)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 128)
}

func NewFNV1a256() Hasher {
	offset, err := hexToBytes(offset256)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime256)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 256)
}
func NewFNV1a512() Hasher {
	offset, err := hexToBytes(offset512)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime512)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 512)
}
func NewFNV1a1024() Hasher {
	offset, err := hexToBytes(offset1024)
	if err != nil {
		panic(err)
	}
	prime, err := hexToBytes(prime1024)
	if err != nil {
		panic(err)
	}
	return newFNV1a(offset, prime, 1024)
}

func newFNV1a(offset []byte, prime []byte, length int) Hasher {
	modulus := new(big.Int).Lsh(big.NewInt(1), uint(length)) // 2^length
	offsetInt := new(big.Int).SetBytes(offset)
	primeInt := new(big.Int).SetBytes(prime)

	return &FNV1a{
		offset: offsetInt,
		prime:  primeInt,
		mod:    modulus,
	}
}

func (h *FNV1a) Hash(data []byte) []byte {
	hash := h.offset
	cont := new(big.Int)

	for _, b := range data {
		c := cont.SetBytes([]byte{b})
		hash.Xor(hash, c)
		hash.Mul(hash, h.prime)
		hash.Mod(hash, h.mod)
	}
	return hash.Bytes()
}
