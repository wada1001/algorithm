package list

import (
	"errors"

	"github.com/wada1001/algorithm/src/pkg/util"
)

// easy to add. but search is difficult
type LinkListNode struct {
	Prev *LinkListNode
	Next *LinkListNode
	Val int
}

func MakeRandomIntLinkListNodes(length int) *LinkListNode {
	top := MakeLinkListNode(nil, util.MakeRandomInt())
	current := top
	i := 0
	for i < length {
		tmp := MakeLinkListNode(nil, util.MakeRandomInt())
		current.Next = tmp
		current = tmp
		i++
	}

	return top
}

func MakeLinkListNode(prev *LinkListNode, val int) *LinkListNode {
	return &LinkListNode{Prev: prev, Val: val}
}


func (l *LinkListNode)GetVal(index int) (int, error) {
	cursor := 0
	tmp := l
	for tmp.Next != nil {
		if cursor == index {
			return tmp.Val, nil
		}
		cursor++
	}
	return 0, errors.New("out of range")
}