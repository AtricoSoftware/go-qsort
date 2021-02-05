package main

import "sync"

func main() {
}

type queue struct {
	queue []int
	len   int
	readMutex sync.Mutex
	writeMutex sync.Mutex
}

func newQueue() queue {
	return queue{queue: make([]int, 0), len: 0}
}

func (q *queue) push(value int) {
	q.writeMutex.Lock()
	defer q.writeMutex.Unlock()
	q.queue = append(q.queue, value)
	q.len++
}

func (q *queue) pop() node {
	// Wait for data available
	q.writeMutex.Lock()
	q.writeMutex.Unlock()

	q.mut.Lock()
	defer q.mut.Unlock()
	if q.isEmpty() {
		panic("Queue is isEmpty")
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return val
}
