package stack_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/stack"
)

func TestPush(t *testing.T) {
	
	type args struct {
		st *stack.Stack
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{&stack.Stack{}, 20}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.st.Push(test.args.val)
			val, _ := test.args.st.Peek()
			if val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, val))
			}
		})
	}
}

func TestPop(t *testing.T) {
	st := &stack.Stack{}
	st.Push(20)
	st.Push(2)
	st.Push(1)
	type args struct {
		st *stack.Stack
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{st, 24}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.st.Push(test.args.val)
			val, err := test.args.st.Pop()
			if err != nil {
				t.Errorf("error occured")
			}

			if val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, val))
			}
		})
	}
}

func TestPeek(t *testing.T) {
	st := &stack.Stack{}
	st.Push(20)
	st.Push(2)
	st.Push(1)
	type args struct {
		st *stack.Stack
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{st, 24}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.st.Push(test.args.val)
			val, err := test.args.st.Peek()
			if err != nil {
				t.Errorf("error occured")
			}

			if val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, val))
			}
		})
	}
}