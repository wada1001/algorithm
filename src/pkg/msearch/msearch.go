package msearch

import (
	"errors"
)

// O (N)
func LinearSearch (arr []int, tar int) (int, error) {
	for i:= 0; i < len(arr); i++ {
		if arr[i] == tar {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

// O (logN) should be sort
func BinarySearch(arr []int, tar int) (int, error) {
	bottom := 0
	top := len(arr)
	for top > bottom{
		mid := (top + bottom) / 2
		if arr[mid] == tar {
			return mid, nil
		}

		if arr[mid] < tar { 
			bottom = mid + 1
		} else {
			top = mid - 1
		}
	}

	return 0, errors.New("not found")
}

// O (logN+n) if exists same val in arr. return smallest index
func BinarySearchLeftEdge(arr []int, tar int) (int, error) {
	bottom := 0
	top := len(arr)
	mid := 0
	for top > bottom{
		mid = (top + bottom) / 2
		if arr[mid] < tar { 
			bottom = mid + 1
		} else {
			top = mid
		}
	}
	if arr[bottom] == tar {
		return bottom, nil
	}
	
	return 0, errors.New("not found")
}

