package msearch_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/msearch"
	"github.com/wada1001/algorithm/src/pkg/msort"
	"github.com/wada1001/algorithm/src/pkg/util"
)

func TestLinearSearch(t *testing.T) {
	arr := util.MakeUnsignedIntArr(100, 100)
	type args struct {
		arr []int
		tar int
		wantErr bool
	}
	tests := []struct {
		name string
		args args
	}{
		
		{"found", args{arr: arr, tar: arr[10], wantErr: false}},
		{"not found", args{arr: arr, tar: -1, wantErr: true}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i, err := msearch.LinearSearch(test.args.arr, test.args.tar)
			if test.args.wantErr && err == nil || !test.args.wantErr && err != nil{
				t.Errorf("error occured")
			}

			if test.args.wantErr {
				return
			}
			
			if test.args.arr[i] != test.args.tar {
				t.Errorf("worng value")
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	arr := util.MakeUnsignedIntArr(100, 100)
	type args struct {
		arr []int
		tar int
		wantErr bool
	}
	tests := []struct {
		name string
		args args
	}{
		
		{"found", args{arr: arr, tar: arr[10], wantErr: false}},
		{"not found", args{arr: arr, tar: -1, wantErr: true}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			msort.CombSort(test.args.arr)
			i, err := msearch.BinarySearch(test.args.arr, test.args.tar)
			if test.args.wantErr && err == nil || !test.args.wantErr && err != nil{
				t.Errorf("error occured")
			}

			if test.args.wantErr {
				return
			}
			
			if test.args.arr[i] != test.args.tar {
				t.Error(arr)
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.tar, test.args.arr[i]))
			}
		})
	}
}