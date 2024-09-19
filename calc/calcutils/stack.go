package calcutils

import "fmt"

type Stack interface {
	Push(val interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	IsEmpty() bool
}

type sliceStack struct {
	elements []interface{}
	top      int
}

func NewStack() *sliceStack {
	return &sliceStack{top: -1}
}

func (s *sliceStack) Push(val interface{}) error {
	if len(s.elements) == cap(s.elements) {
		s.elements = append(s.elements, val)
		s.top++
		return nil
	}

	s.elements[s.top] = val
	s.top++
	return nil
}

func (s *sliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, nil
	}

	top := s.top
	s.top--
	return s.elements[top], nil
}

func (s *sliceStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.elements[s.top], nil
}

func (s *sliceStack) IsEmpty() bool {
	return s.top < 0
}
