package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cases map[string][]int

func main() {
	cases := make(cases)
	addList(&cases, 30)
	// addList(&cases, 100)
	// addList(&cases, 10000)
	// addList(&cases, 1000000)

	for k, v := range cases {
		fmt.Println(k)
		if len(v) <= 100 {
			fmt.Println(v)
		}
		start := time.Now()
		Qsort(v)
		elapsed := time.Now().Sub(start)
		fmt.Printf("Time taken: %s\n", elapsed)
		if len(v) <= 100 {
			fmt.Println(v)
		}
	}
}

func addList(container *cases, size int) {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(100)
	}
	(*container)[string(size)] = list
}
