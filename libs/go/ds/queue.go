package ds

import (
	"errors"
)

type Queue struct {
	backingStore []interface{} // default is empty
}

func (q *Queue) Enqueue(toEnqueue interface{}) {
	q.backingStore = append(q.backingStore, toEnqueue)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	result := q.backingStore[0]
	q.backingStore = q.backingStore[1:]
	return result, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.backingStore) == 0
}
