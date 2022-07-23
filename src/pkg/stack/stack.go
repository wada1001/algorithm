package stack

import (
	"errors"

	"github.com/wada1001/algorithm/src/pkg/list"
)

type Stack struct {
	node *list.LinkListNode
}


func (s *Stack) Push(val int) {
	n := list.MakeLinkListNode(nil, val)
	if s.node == nil {
		s.node = n
		return
	}
	
	s.node.Next = n
	s.node = n
}

func (s *Stack) Pop() (int, error) {
	if s.node == nil {
		return 0, errors.New("cannot pop")
	}
	
	v := s.node.Val
	s.node = s.node.Prev
	return v, nil
}

func (s *Stack) Peek() (int, error) {
	if s.node == nil {
		return 0, errors.New("cannot peek")
	}

	return s.node.Val, nil
}
