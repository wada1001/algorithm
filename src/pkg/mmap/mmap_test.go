package mmap_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/mmap"
	"github.com/wada1001/algorithm/src/pkg/util"
)

// TODO test HashMao
func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"gen"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := make([]string, 100)
			i := 0
			for i * 2 < 100 {
				str := util.MakeRandomStr(5)
				tmp := mmap.MakeHash(str, 100)
				arr[tmp] = str
				fmt.Print(str + " ")
				fmt.Println(tmp)
				i++
			}
		})
	}
}