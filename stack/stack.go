package stack

type Stack[T any] struct {
	length int
	head   *Node[T]
}

type Node[T any] struct {
	prev  *Node[T]
	value T
}

func (s *Stack[T]) push(v T) {
	node := Node[T]{
		value: v,
	}
	if s.length == 0 {
		s.head = &node
	}
	s.length++

	node.prev = s.head
	s.head = &node
}

func (s *Stack[T]) pop() T {
	var n Node[T]
	if s.length != 0 {
		head := s.head
		s.head = head.prev
		return head.value
	} else {
		head := s.head
		s.head = &n
		return head.value
	}
}

func (s *Stack[T]) peek() T {
	return s.head.value
}
