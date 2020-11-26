package main

func Qsort(list []int) {
	qsort(list, 0, len(list)-1)
}

func qsort(list []int, start int, end int) {
	if end <= start {
		return
	}
	pivot := (list[start] + list[end]) / 2
	left := start
	right := end
	for left < right {
		for left < right && list[left] < pivot {
			left++
		}
		for right > left && list[right] > pivot {
			right--
		}
		if left < right {
			swap(list, left, right)
		}
	}
	startL := start
	endL := left
	startR := right
	endR:=end
	if list[left] > pivot {
		endL--
	} else {
		startR++
	}
	qsort(list, startL, endL)
	qsort(list, startR, endR)
}

func swap(list []int, idx1 int, idx2 int) {
	tmp := list[idx1]
	list[idx1] = list[idx2]
	list[idx2] = tmp
}
