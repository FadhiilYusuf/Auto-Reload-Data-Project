package helper

import (
	"crypto/rand"
	"math/big"
)

func RandomNumber(min, max int64) int64 {
	number := big.NewInt(max - min)

	num, err := rand.Int(rand.Reader, number)
	if err != nil {
		panic(err)
	}

	return num.Int64() + min
}
