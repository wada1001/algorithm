package msort_test

import (
	"testing"

	"github.com/wada1001/algorithm/src/pkg/msort"
	"github.com/wada1001/algorithm/src/pkg/prepare"
)


func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_a", args{arr: prepare.MakeUnsignedIntArr(100, 100)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := test.args.arr
			msort.BubbleSort(arr)
			for i := 1; i < len(arr); i++ {
				if arr[i - 1] > arr[i] {
					t.Errorf("Invalid value")
				}
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_a", args{arr: prepare.MakeUnsignedIntArr(100, 100)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := test.args.arr
			msort.QuickSort(0, len(arr) - 1, arr)
			for i := 1; i < len(arr); i++ {
				if arr[i - 1] > arr[i] {
					t.Errorf("Invalid value")
				}
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_a", args{arr: prepare.MakeUnsignedIntArr(100, 100)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := test.args.arr
			msort.MergeSort(arr, make([]int, len(arr)))
			for i := 1; i < len(arr); i++ {
				if arr[i - 1] > arr[i] {
					t.Errorf("Invalid value")
				}
			}
		})
	}
}

func TestCombSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_a", args{arr: prepare.MakeUnsignedIntArr(100, 100)}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := test.args.arr
			msort.CombSort(arr)
			for i := 1; i < len(arr); i++ {
				if arr[i - 1] > arr[i] {
					t.Errorf("Invalid value")
				}
			}
		})
	}
}