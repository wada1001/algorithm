package txt_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/txt"
)


func TestKmp(t *testing.T) {
	tests := []struct {
		name string
		str string
		search string
		want int
	}{
		{"found", "kontinaoksuenfoodgon", "ksuen", 8},
		{"not found", "kontinaoksuenfoodgon", "inaokd", -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index := txt.Kmp(test.str, test.search)
			if index != test.want {
				t.Errorf(fmt.Sprintf("worng value got = %d", index))
			}
		})
	}
}

func TestBm(t *testing.T) {
	tests := []struct {
		name string
		str string
		search string
		want int
	}{
		// TODO more case
		{"found", "kontinaoksuenfoodgon", "ksuen", 8},
		{"found", "aiueokeksuen", "ksuen", 7},
		{"not found", "kontinaoksunfoodgon", "ksuen", -1},
		{"not found", "kontinaoksuenfoodgon", "inaokd", -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index := txt.Bm(test.str, test.search)
			if index != test.want {
				t.Errorf(fmt.Sprintf("worng value got = %d", index))
			}
		})
	}
}