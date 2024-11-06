package stack

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	top *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.top == nil {
		var zeroT T
		return zeroT, false
	}

	return s.top.value, true
}

func (s *Stack[T]) Push(v T) {
	s.top = &Node[T]{
		value: v,
		next:  s.top,
	}
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var zeroT T
		return zeroT, false
	}

	value := s.top.value
	s.top = s.top.next
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) GetItems() []T {
	var values []T
	current := s.top

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}
