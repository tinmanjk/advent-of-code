package ds

import (
	"errors"
)

type Stack []rune

func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (result rune, err error) {

	if s.IsEmpty() {
		return result, errors.New("empty stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *Stack) IsEmpty() bool {

	return len(*s) < 1
}
