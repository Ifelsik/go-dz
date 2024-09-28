package calc

type Stack interface {
	Push(val interface{}) bool
	Pop() (interface{}, bool)
	Top() (interface{}, bool)
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

func (s *sliceStack) Push(val string) bool {
	if len(s.elements) == s.size {
		s.size *= 2
		tmp := s.elements
		s.elements = make([]string, s.size)
		copy(s.elements, tmp)
	}

	s.top++
	s.elements[s.top] = val
	return true
}

func (s *sliceStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	top := s.top
	s.top--
	return s.elements[top], true
}

func (s *sliceStack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	return s.elements[s.top], true
}

func (s *sliceStack) IsEmpty() bool {
	return s.top < 0
}
