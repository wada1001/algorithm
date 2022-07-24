package tree_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/tree"
)

func TestMakeNode(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{val: 100}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node := tree.MakeNode(test.args.val)
			if node.Val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, node.Val))
			}
		})
	}
}

func TestInsertNode(t *testing.T) {
	nodet1 := tree.MakeNode(100)
	nodet1.Left = tree.MakeNode(59)
	nodet1.Right = tree.MakeNode(201)

	// should Deep Copy
	nodet2 := tree.MakeNode(100)
	nodet2.Left = tree.MakeNode(59)
	nodet2.Right = tree.MakeNode(201)

	node2 := tree.MakeNode(100)
	node2.Left = tree.MakeNode(59)
	node2.Right = tree.MakeNode(201)
	node2.Right.Right = tree.MakeNode(205)

	node3 := tree.MakeNode(100)
	node3.Left = tree.MakeNode(59)
	node3.Right = tree.MakeNode(201)
	node3.Left.Left = tree.MakeNode(49)
	
	type args struct {
		node *tree.TreeNode
		want *tree.TreeNode
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_0", args{node: nil, want: tree.MakeNode(205), val: 205}},
		{"test_1", args{node: nodet1, want: node2, val: 205}},
		{"test_2", args{node: nodet2, want: node3, val: 49}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node := tree.InsertNode(test.args.val, test.args.node)
			if !reflect.DeepEqual(node, test.args.want) {
				t.Errorf(fmt.Sprintf("worng value want = %+v, got = %+v" , test.args.want, node))
			}
		})
	}
}

func TestSearchNode(t *testing.T) {
	nodet1 := tree.MakeNode(100)
	nodet1.Left = tree.MakeNode(59)
	nodet1.Left.Left = tree.MakeNode(49)
	nodet1.Left.Right = tree.MakeNode(99)
	nodet1.Right = tree.MakeNode(201)
	nodet1.Right.Left = tree.MakeNode(101)
	nodet1.Right.Right = tree.MakeNode(271)

	type args struct {
		node *tree.TreeNode
		tar int
		want bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"success_1", args{node: nodet1, tar: 100, want: true}},
		{"success_2", args{node: nodet1, tar: 99, want: true}},
		{"success_3", args{node: nodet1, tar: 271, want: true}},
		{"success_4", args{node: nodet1, tar: 101, want: true}},
		{"not found", args{node: nodet1, tar: 10, want: false}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node := tree.SearchNode(test.args.tar, test.args.node)
			if test.args.want && node == nil {
				t.Errorf("should found")
			}

			if !test.args.want {
				return
			}

			if node.Val != test.args.tar {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.tar, node.Val))
			}
		})
	}
}


func TestSearchBiggest(t *testing.T) {
	nodet1 := tree.MakeNode(100)
	nodet1.Left = tree.MakeNode(59)
	nodet1.Left.Left = tree.MakeNode(49)
	nodet1.Left.Right = tree.MakeNode(99)
	nodet1.Right = tree.MakeNode(201)
	nodet1.Right.Left = tree.MakeNode(101)
	nodet1.Right.Right = tree.MakeNode(271)

	type args struct {
		node *tree.TreeNode
		tar int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success_1", args{node: nodet1, tar: 271}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node := tree.SearchBiggest(test.args.node)
			if node.Val != test.args.tar {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.tar, node.Val))
			}
		})
	}
}

// TODO Test
// func TestDeleteNode(t *testing.T) {
// 	nodet1 := tree.MakeNode(100)
// 	nodet1.Left = tree.MakeNode(59)
// 	nodet1.Left.Left = tree.MakeNode(49)
// 	nodet1.Left.Right = tree.MakeNode(99)
// 	nodet1.Right = tree.MakeNode(201)
// 	nodet1.Right.Left = tree.MakeNode(101)
// 	nodet1.Right.Right = tree.MakeNode(271)

// 	type args struct {
// 		node *tree.TreeNode
// 		tar int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{"success_1", args{node: nodet1, tar: 271}},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			node := tree.SearchBiggest(test.args.node)
// 			if node.Val != test.args.tar {
// 				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.tar, node.Val))
// 			}
// 		})
// 	}
// }


