package main

import (
	"fmt"

	"github.com/wada1001/algorithm/src/pkg/stopwatch"
	"github.com/wada1001/algorithm/src/pkg/tree"
)

func main ()  {
	w := stopwatch.Make()
	w.Start()
	w.Lap()
	node1 := tree.MakeNode(100)
	node1.Left = tree.MakeNode(59)
	node1.Right = tree.MakeNode(201)

	node := tree.InsertNode(205, node1)
	fmt.Println(node)
	w.Lap()
}
