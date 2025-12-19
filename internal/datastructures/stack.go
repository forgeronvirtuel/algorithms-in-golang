package datastructures

type nodeStack[T any] struct {
	value T
	next  *nodeStack[T]
}

type Stack[T any] struct {
	head *nodeStack[T]
	size int
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{head: nil, size: 0}
}

// Push adds a new element to the top of the stack
// Complexity: O(1)
func (s *Stack[T]) Push(value T) {
	newNode := &nodeStack[T]{value: value, next: s.head}
	s.head = newNode
	s.size++
}

// Pop removes and returns the top element of the stack
// Complexity: O(1)
func (s *Stack[T]) Pop() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	topNode := s.head
	s.head = topNode.next
	s.size--
	return topNode.value, true
}

// Peek returns the top element of the stack without removing it
// Complexity: O(1)
func (s *Stack[T]) Peek() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}
	return s.head.value, true
}

// Size returns the number of elements in the stack
// Complexity: O(1)
func (s *Stack[T]) Size() int {
	return s.size
}

// IsEmpty checks if the stack is empty
// Complexity: O(1)
func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}
