package datastructures

type nodeQueue[T any] struct {
	value T
	next  *nodeQueue[T]
}

type Queue[T any] struct {
	head *nodeQueue[T]
	tail *nodeQueue[T]
	size int
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{head: nil, tail: nil, size: 0}
}

// Enqueue adds a new element to the end of the queue
// Complexity: O(1)
func (q *Queue[T]) Enqueue(value T) {
	newNode := &nodeQueue[T]{value: value, next: nil}
	if q.tail != nil {
		q.tail.next = newNode
	} else {
		q.head = newNode
	}
	q.tail = newNode
	q.size++
}

// Dequeue removes and returns the front element of the queue
// Complexity: O(1)
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	frontNode := q.head
	q.head = frontNode.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return frontNode.value, true
}

// Peek returns the front element of the queue without removing it
// Complexity: O(1)
func (q *Queue[T]) Peek() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.head.value, true
}

// Size returns the number of elements in the queue
// Complexity: O(1)
func (q *Queue[T]) Size() int {
	return q.size
}

// IsEmpty checks if the queue is empty
// Complexity: O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}
