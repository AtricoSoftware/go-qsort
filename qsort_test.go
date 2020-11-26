package main

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func TestEmpty(t *testing.T) {
	// Arrange
	list := make([]int, 0)

	// Act
	fmt.Println(list)
	Qsort(list)
	fmt.Println(list)

	// Assert
	// No panic!
}

func TestSingle(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value}

	// Act
	fmt.Println(list)
	Qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(list[0], is.EqualTo(value), "Same value")
}

func TestTwoInOrder(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value, value + 1}

	// Act
	fmt.Println(list)
	Qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(list[0], is.EqualTo(value), "Same value 1")
	Assert(t).That(list[1], is.EqualTo(value+1), "Same value 2")
}

func TestTwoOutOfOrder(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value + 1, value}

	// Act
	fmt.Println(list)
	Qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(list[0], is.EqualTo(value), "Same value 1")
	Assert(t).That(list[1], is.EqualTo(value+1), "Same value 2")
}

func TestThree(t *testing.T) {
	// Arrange
	list := createList(3)

	// Act
	fmt.Println(list)
	Qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(isSorted(list), is.True, "Sorted")
}

func createList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(100)
	}
	return list
}

func isSorted(list []int) bool {
	if len(list) < 2 {
		return true
	}
	val := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < val {return false}
		val = list[i]
	}
	return true;
}
