package main

import (
	"crypto/md5"
	"math/big"
)

type Hasher interface {
	CalculateHash(key string) *big.Int
	Distance(a, b *big.Int) *big.Int
}
type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

func (h *HashService) CalculateHash(key string) *big.Int {
	hash := md5.Sum([]byte(key))
	hashInt := new(big.Int).SetBytes(hash[:])
	return hashInt
}
func (h *HashService) Distance(a *big.Int, b *big.Int) *big.Int {

	return new(big.Int).Abs(new(big.Int).Sub(a, b))
}

// func main() {
// 	// hash := md5.New()
// 	// hash.Sum()
// }
