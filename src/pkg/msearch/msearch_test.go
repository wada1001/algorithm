package msearch_test

import (
	"testing"

	"github.com/wada1001/algorithm/src/pkg/msearch"
	"github.com/wada1001/algorithm/src/pkg/prepare"
)

func TestCombSort(t *testing.T) {
	arr := prepare.MakeUnsignedIntArr(100, 100)
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
			if test.args.wantErr && err == nil {
				t.Errorf("should be error")
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