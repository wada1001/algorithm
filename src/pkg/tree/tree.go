package tree

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func MakeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func InsertNode(val int, node *TreeNode) *TreeNode {
	if node == nil {
		return MakeNode(val)
	}
	
	if val < node.Val {
		if node.Left != nil {
			InsertNode(val, node.Left)
		} else {
			fmt.Print("set to left = ")
			fmt.Println(val)
			
			node.Left = MakeNode(val)
		}
	} else {
		if node.Right != nil {
			InsertNode(val, node.Right)
		} else {
			node.Right = MakeNode(val)
		}
	}
	return node
}

func SearchNode(val int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	
	if node.Val == val {
		return node
	}

	if val < node.Val {
		return SearchNode(val, node.Left)
	} else {
		return SearchNode(val, node.Right)
	}
}

func DeleteNode(val int, node *TreeNode) bool {
	if node == nil {
		return false
	}

	parent := node
	current := node
	for current != nil && current.Val != val {
		if val < current.Val {
			parent = current
			current = current.Left
		} else {
			parent = current
			current = current.Right
		}
	}
	
	if current == nil {
		return false
	}

	if current.Left == nil && current.Right == nil {
		if parent.Left.Val == current.Val {
			parent.Left = nil
		}
		if parent.Right.Val == current.Val {
			parent.Right = nil
		}
	} else if current.Left == nil {
		if parent.Left.Val == current.Val {
			parent.Left = current.Right
		}
		if parent.Right.Val == current.Val {
			parent.Right = current.Right
		}
	} else if current.Right == nil {
		if parent.Left.Val == current.Val {
			parent.Left = current.Left
		}
		if parent.Right.Val == current.Val {
			parent.Right = current.Left
		}
	} else {
		// LR exists
		desparent := current.Left
		descurrent := current.Left
		for current.Right != nil {
			desparent = descurrent
			descurrent = current.Right
		}
		// free node
		desparent.Right = nil
		// swap current and descurrent
		descurrent.Left = current.Left
		descurrent.Right = current.Right

		if parent.Left.Val == current.Val {
			parent.Left = descurrent
		}
		if parent.Right.Val == current.Val {
			parent.Right = descurrent
		}
	}
	return true
}

// find biggest descendant in node
func SearchBiggest(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return node
	}

	return SearchBiggest(node.Right)
}