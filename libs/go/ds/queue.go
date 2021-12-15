package ds

import (
	"errors"
)

type MyQueue struct {
	backingStore []int // default is empty
}

func (q *MyQueue) Enqueue(toEnqueue int) {
	q.backingStore = append(q.backingStore, toEnqueue)
}

func (q *MyQueue) Dequeue() (int, error) {
	if q.IsEmtpy() {
		return -1, errors.New("empty queue")
	}
	result := q.backingStore[0]
	q.backingStore = q.backingStore[1:]
	return result, nil
}

func (q *MyQueue) IsEmtpy() bool {
	return len(q.backingStore) == 0
}
