package main

func Qsort(list []int) {
	qsort(list, 0, len(list)-1)
}

func qsort(list []int, start int, end int) {
	if end <= start {
		return
	}
	pivot := list[end]
	j := start - 1
	for i := start; i < end; i++ {
		if list[i] < pivot {
			j++
			swap(list, i, j)
		}
	}
	swap(list, j+1, end)

	qsort(list, start, j)
	qsort(list, j+2, end)
}

func swap(list []int, idx1 int, idx2 int) {
	tmp := list[idx1]
	list[idx1] = list[idx2]
	list[idx2] = tmp
}
