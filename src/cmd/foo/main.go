package main

import (
	"github.com/wada1001/algorithm/src/pkg/recursive"
	"github.com/wada1001/algorithm/src/pkg/stopwatch"
)

func main ()  {
	w := stopwatch.Make()
	w.Start()
	w.Lap()
	p := recursive.MakePuzzle(3, 3)
	p.Solve()
	w.Lap()
}
