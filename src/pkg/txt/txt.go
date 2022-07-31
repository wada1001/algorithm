package txt

import (
	"strings"
)

func Kmp(str, search string) int {
	index := -1
	for i := 0; i < len(str); i++ {
		if str[i] != search[0] {
			continue
		}

		if len(str) - i < len(search) {
			break
		}

		for j := 0; j < len(search); j++ {
			if str[i+j] != search[j] {
				// skip j
				i += j
				break
			}

			if j == len(search) - 1 {
				return i
			} 
		}
	}
	return index
}

func Bm(str, search string) int {
	index := -1
	searchLength := len(search) - 1
	for i := searchLength; i < len(str); i++ {
		if str[i] != search[searchLength] {
			tmpIndex := strings.Index(search, string(str[i]))
			if tmpIndex >= 0 {
				i += searchLength - tmpIndex - 1
			} else{
				i += searchLength
			}
			continue
		}
		
		tmp := 0
		for j := searchLength; j >= 0; j-- {
			if str[i-tmp] != search[j] {
				i += searchLength
				break
			}

			if j == 0 {
				return i - searchLength
			}
			tmp++;
		}
	}
	return index
}