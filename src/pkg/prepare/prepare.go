package prepare

import (
	"math/rand"
	"time"
)

// const seed = rand.Seed(time.Now().UnixNano())

// generate not negative value arr
func MakeUnsignedIntArr(size, max int) []int {
	rand.Seed(time.Now().UnixNano())
	i := 0
	arr := make([]int, size)
	for i < size {
		arr[i] = rand.Intn(max)
		i++
	}
	return arr
}

// generate include negative value arr
func MakeIntArr(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	i := 0
	arr := make([]int, size)
	for i < size {
		arr[i] = rand.Intn(max - min) + min
		i++
	}
	return arr
}