package main

import "fmt"

func Qsort(list []int) {
	if len(list) < 2 {
		return
	}
	qsort(list, 0, len(list)-1, 0)
}

func qsort(list []int, start int, end int, indent int) {
	indentedPrint(indent, list, start, end)
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
			swap(list, left, right, indent)
		}
	}
	startL := start
	endL := left
	startR := right
	endR := end
	if list[left] > pivot {
		endL--
	} else {
		startR++
	}
	if (endL - startL) > 0 {qsort(list, startL, endL, indent+1)}
	if (endR - startR) > 0 {qsort(list, startR, endR, indent+1)}
	indentedPrint(indent, list)
}

func swap(list []int, idx1 int, idx2 int, indent int) {
	indentedPrint(indent, fmt.Sprintf("sw(%v,%v)", idx1, idx2))
	tmp := list[idx1]
	list[idx1] = list[idx2]
	list[idx2] = tmp
}

func indentedPrint(indent int, args ...interface{}) {
	for i:=0; i< indent;i++ {
		fmt.Print(" ")
	}
	for _,arg := range args {
		fmt.Printf("%v, ", arg)
	}
	fmt.Println()
}