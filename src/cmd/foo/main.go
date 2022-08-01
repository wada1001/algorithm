package main

import (
	"github.com/wada1001/algorithm/src/pkg/dynamic"
	"github.com/wada1001/algorithm/src/pkg/stopwatch"
)

func main ()  {
	w := stopwatch.Make()
	w.Start()
	w.Lap()
	dynamic.LessCoinCountPay()
	w.Lap()
}
