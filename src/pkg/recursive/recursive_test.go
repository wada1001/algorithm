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