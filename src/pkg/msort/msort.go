package msort

// O (N^2)
func BubbleSort(arr []int) {
	count := len(arr)
	modify := true
	tmp := 0
	for modify {
		modify = false
		for i := 1; i < count; i++ {
			if arr[i - 1] <= arr[i] {
				continue;
			}
			tmp = arr[i - 1]
			arr[i - 1] = arr[i]
			arr[i] = tmp
			modify = true
		}
	}
}

// O (N^2) ~ (N*logN)
func QuickSort(bottom, top int, arr []int) {
	if bottom >= top {
		return
	}
	
	low := bottom
	hight := top
	tar := arr[bottom]
	for low < hight {
		for low <= hight && tar >= arr[low] {
			low++
		}
		for low <= hight && tar < arr[hight] {
			hight--
		}
		if low < hight {
			// swap tar >= arr[low], tar < arr[hight]
			tmp := arr[hight]
			arr[hight] = arr[low]
			arr[low] = tmp
		}
	}

	// swap tar, arr[hight]
	arr[bottom] = arr[hight]
	arr[hight] = tar
	QuickSort(bottom, hight - 1, arr)
	QuickSort(hight + 1, top, arr)
}

// O (N*logN) stable
func MergeSort (arr []int, buff []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	
	mid := length / 2
	MergeSort(arr[0:mid], buff)
	MergeSort(arr[mid:length], buff)

	// buff set arr[0:mid]
	for i := 0; i < mid; i++ {
		buff[i] = arr[i]
	}

	q := mid
	s, t := 0, 0
	for s < mid && q < length {
		// compare buff, arr[mid + 1:length]
		if buff[s] <= arr[q] {
			// original arr[t] = buff[s]
			arr[t] = buff[s]
			s++
			t++
		} else {
			arr[t] = arr[q]
			q++
			t++
		}
	}
	for s < mid {
		arr[t] = buff[s]
		s++
		t++
	}
}

// O (N*logN)
func CombSort(arr []int) {
	length := len(arr)
	gap := length

	done := false
 	for !done || gap > 1{
		// 1.3 approaching the optimum
		gap = (gap * 10) / 13
		if gap == 0 {
			gap = 1
		}
		done = true
		for i := 0; i + gap < length; i++ {
			if arr[i] <= arr[i + gap] {
				continue
			}
			tmp := arr[i]
			arr[i] = arr[i+gap]
			arr[i + gap] = tmp
			done = false
		}
	}
}

// O (N^2)
func InsertSort(arr []int) {
	for i:= 1; i < len(arr); i++ {
		for j:= i; j >= 1; j-- {
			if arr[j] >= arr[j - 1] {
				continue
			}
			tmp := arr[j]
			arr[j] = arr[j - 1]
			arr[j - 1] = tmp
		}	
	}
}
