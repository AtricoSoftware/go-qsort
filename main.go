package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

var repeat = 8
var size = int(math.Pow(2,20))

func main() {
	var total time.Duration = 0
	var total2 time.Duration = 0
	for i := 0; i < repeat; i++ {
		list,list2 := makeList(size)
		start := time.Now()
		Qsort(list)
		elapsed := time.Now().Sub(start)
		start = time.Now()
		sort.Ints(list2)
		elapsed2 := time.Now().Sub(start)
		fmt.Printf("%s\t%s\n", elapsed, elapsed2)
		total += elapsed
		total2 += elapsed2
	}
	fmt.Printf("Total  : %s\t%s\n", total, total2)
	fmt.Printf("Average: %s\t%s\n", total/time.Duration(repeat), total2/time.Duration(repeat))
}

func makeList(size int) ([]int,[]int) {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(100)
	}
	list2 := make([]int,size)
	copy(list2,list)
	return list,list2
}
