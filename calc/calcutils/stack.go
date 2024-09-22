package calcutils

import "fmt"

type Stack interface {
	Push(val interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	IsEmpty() bool
}

type sliceStack struct {
	elements []string
	top      int
	size	 int
}

func NewStack() *sliceStack {
	return &sliceStack{elements: make([]string, 1), top: -1, size: 1}
}

func (s *sliceStack) Push(val string) error {
	if len(s.elements) == s.size {
		s.size *= 2
		tmp := s.elements
		s.elements = make([]string, s.size)
		copy(s.elements, tmp)
	}

	s.top++
	s.elements[s.top] = val
	return nil
}

func (s *sliceStack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}

	top := s.top
	s.top--
	return s.elements[top], nil
}

func (s *sliceStack) Top() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	return s.elements[s.top], nil
}

func (s *sliceStack) IsEmpty() bool {
	return s.top < 0
}
