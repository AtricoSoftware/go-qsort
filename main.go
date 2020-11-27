package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var repeat = 8
var size = int(math.Pow(2,20))

func main() {
	var total time.Duration = 0
	for i := 0; i < repeat; i++ {
		list := makeList(size)

		start := time.Now()
		Qsort(list)
		elapsed := time.Now().Sub(start)
		fmt.Printf("%s\n", elapsed)
		total += elapsed
	}
	fmt.Printf("Total  : %s\n", total)
	fmt.Printf("Average: %s\n", total/time.Duration(repeat))
}

func makeList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(100)
	}
	return list
}
