package main

import "sync"

func Qsort(list []int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go qsort(list, 0, len(list)-1, &wg)
	wg.Wait()
}

func qsort(list []int, start int, end int, wg *sync.WaitGroup) {
	if end > start {
		pivot := list[end]
		j := start - 1
		for i := start; i < end; i++ {
			if list[i] < pivot {
				j++
				swap(list, i, j)
			}
		}
		swap(list, j+1, end)

		wg2 := sync.WaitGroup{}
		wg2.Add(2)
		go qsort(list, start, j, &wg2)
		go qsort(list, j+2, end, &wg2)
		wg2.Wait()
	}
	wg.Done()
}

func swap(list []int, idx1 int, idx2 int) {
	tmp := list[idx1]
	list[idx1] = list[idx2]
	list[idx2] = tmp
}
