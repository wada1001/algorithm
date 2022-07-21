package msearch

import "errors"

func LinearSearch (arr []int, tar int) (int, error) {
	for i:= 0; i < len(arr); i++ {
		if arr[i] == tar {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}