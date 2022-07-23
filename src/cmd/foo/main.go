package main

import (
	"fmt"

	"github.com/wada1001/algorithm/src/pkg/msearch"
	"github.com/wada1001/algorithm/src/pkg/msort"
	"github.com/wada1001/algorithm/src/pkg/prepare"
	"github.com/wada1001/algorithm/src/pkg/stopwatch"
)

func main ()  {
	w := stopwatch.Make()
	w.Start()
	w.Lap()
	arr := prepare.MakeUnsignedIntArr(10, 100)
	arr = []int { 1, 2, 2, 2, 2, 2, 3, 4, 4, 4, 5}
	msort.BubbleSort(arr)
	fmt.Println(arr)
	i, err := msearch.BinarySearchLeftEdge(arr, 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)
	w.Lap()
}

func sampleFuncGenerics1[T comparable](x T) T {
	return x
}