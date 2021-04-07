package util

import (
	"math/rand"
)

//从0-(maxval-1)中随机挑选quantity个序号
func UniqRands(quantity int, maxval int) []int {
	if maxval < quantity {
		quantity = maxval
	}

	intSlice := make([]int, maxval)
	for i := 0; i < maxval; i++ {
		intSlice[i] = i
	}

	for i := 0; i < quantity; i++ {
		j := rand.Int()%maxval + i
		// swap
		intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
		maxval--

	}
	return intSlice[0:quantity]
}
