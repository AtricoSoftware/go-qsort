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
	qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(len(list), is.EqualTo(0), "Empty result")
}

func TestSingle(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value}

	// Act
	fmt.Println(list)
	qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(len(list), is.EqualTo(1), "Single result")
	Assert(t).That(list[0], is.EqualTo(value), "Same value")
}

func TestTwoInOrder(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value, value + 1}

	// Act
	fmt.Println(list)
	qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(len(list), is.EqualTo(2), "Same size")
	Assert(t).That(list[0], is.EqualTo(value), "Same value 1")
	Assert(t).That(list[1], is.EqualTo(value+1), "Same value 2")
}

func TestTwoOutOfOrder(t *testing.T) {
	// Arrange
	value := rand.Intn(100)
	list := []int{value + 1, value}

	// Act
	fmt.Println(list)
	qsort(list)
	fmt.Println(list)

	// Assert
	Assert(t).That(len(list), is.EqualTo(2), "Same size")
	Assert(t).That(list[0], is.EqualTo(value), "Same value 1")
	Assert(t).That(list[1], is.EqualTo(value+1), "Same value 2")
}

