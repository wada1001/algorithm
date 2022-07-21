package main

import (
	"fmt"

	"github.com/wada1001/algorithm/src/pkg/msearch"
	"github.com/wada1001/algorithm/src/pkg/prepare"
	"github.com/wada1001/algorithm/src/pkg/stopwatch"
)

func main ()  {
	w := stopwatch.Make()
	w.Start()
	w.Lap()
	arr := prepare.MakeUnsignedIntArr(10, 100)
	arr = []int { 57, 62, 1, 5, 39, 8, 9, 20}
	i, err := msearch.LinearSearch(arr, -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	w.Lap()
}

func Do() {
	
}