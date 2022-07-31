package recursive_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/recursive"
)


func TestCountOneRecursive(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{val: 11315910120}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := recursive.CountOneRecursive(test.args.val)
			if count != 5 {
				t.Errorf(fmt.Sprintf("worng value want = 5, got = %d" , count))
			}
		})
	}
}

func TestCountOneFor(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{val: 11315910120}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := recursive.CountOneFor(test.args.val)
			if count != 5 {
				t.Errorf(fmt.Sprintf("worng value want = 5, got = %d" , count))
			}
		})
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{nums: []int{60, 40, 900}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := recursive.Gcd(test.args.nums)
			// count := recursive.GcdIn(200, 15)
			if count != 20 {
				t.Errorf(fmt.Sprintf("worng value want = 20, got = %d" , count))
			}
		})
	}
}

func TestCheckQueen(t *testing.T) {
	board := [8][8]int {
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	type args struct {
		board [8][8]int
		x int
		y int
		want bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"ok", args{board: board, x: 2, y: 4, want: false}},
		{"ng", args{board: board, x: 3, y: 2, want: true}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := recursive.CheckQueen(test.args.board, test.args.x, test.args.y)
			if test.args.want != res{
				t.Errorf("error occured")
			}
		})
	}
}

func TestEightQueen(t *testing.T) {
	type args struct {
		want bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{want: true}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := [8][8]int{}
			res := recursive.EightQueen(b, 0, 0)
			if test.args.want != res{
				t.Errorf("error occured")
			}
		})
	}
}