package ds

import (
	"errors"
)

type Stack struct {
	backingStore []interface{} // default is empty
}

func (s *Stack) Push(toEnqueue interface{}) {
	s.backingStore = append(s.backingStore, toEnqueue)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	result := s.backingStore[len(s.backingStore)-1]
	s.backingStore = s.backingStore[:len(s.backingStore)-1]
	return result, nil
}

func (q *Stack) IsEmpty() bool {
	return len(q.backingStore) == 0
}
