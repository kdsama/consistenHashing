package main

import (
	"math/big"
)

type Hasher interface {
	CalculateHash(key string) []byte
	Distance(a, b []byte) big.Int
}

// func main() {
// 	// hash := md5.New()
// 	// hash.Sum()
// }
