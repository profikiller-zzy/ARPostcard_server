package utils

import (
	"math/rand"
	"time"
)

func RandomInt64() int64 {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	min := int64(1000000000000000)   // Minimum 16-digit number
	max := int64(9999999999999999)   // Maximum 16-digit number
	return rand.Int63n(max-min+1) + min
}
