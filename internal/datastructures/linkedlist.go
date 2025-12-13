package datastructures

import "fmt"

// Node represents a node in the linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a singly linked list
type LinkedList[T comparable] struct {
	Head *Node[T]
	size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		size: 0,
	}
}

// InsertAtBeginning adds a new node at the beginning of the list
// Complexity: O(1)
func (ll *LinkedList[T]) InsertAtBeginning(value T) {
	newNode := &Node[T]{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.size++
}

// InsertAtEnd adds a new node at the end of the list
// Complexity: O(n)
func (ll *LinkedList[T]) InsertAtEnd(value T) {
	newNode := &Node[T]{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.size++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.size++
}

// InsertAtPosition inserts a node at a specific position (0-indexed)
// Complexity: O(n)
func (ll *LinkedList[T]) InsertAtPosition(value T, position int) error {
	if position < 0 || position > ll.size {
		return fmt.Errorf("invalid position: %d", position)
	}

	if position == 0 {
		ll.InsertAtBeginning(value)
		return nil
	}

	newNode := &Node[T]{Value: value}
	current := ll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.size++
	return nil
}

// Delete removes the first occurrence of a value from the list
// Complexity: O(n)
func (ll *LinkedList[T]) Delete(value T) bool {
	if ll.Head == nil {
		return false
	}

	// If head needs to be deleted
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}

	return false
}

// DeleteAtPosition removes a node at a specific position (0-indexed)
// Complexity: O(n)
func (ll *LinkedList[T]) DeleteAtPosition(position int) error {
	if position < 0 || position >= ll.size {
		return fmt.Errorf("invalid position: %d", position)
	}

	if position == 0 {
		ll.Head = ll.Head.Next
		ll.size--
		return nil
	}

	current := ll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	ll.size--
	return nil
}

// Search finds if a value exists in the list
// Complexity: O(n)
func (ll *LinkedList[T]) Search(value T) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Get returns the value at a specific position (0-indexed)
// Complexity: O(n)
func (ll *LinkedList[T]) Get(position int) (T, error) {
	var zero T
	if position < 0 || position >= ll.size {
		return zero, fmt.Errorf("invalid position: %d", position)
	}

	current := ll.Head
	for i := 0; i < position; i++ {
		current = current.Next
	}
	return current.Value, nil
}

// Size returns the number of elements in the list
// Complexity: O(1)
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty checks if the list is empty
// Complexity: O(1)
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

// Clear removes all elements from the list
// Complexity: O(1)
func (ll *LinkedList[T]) Clear() {
	ll.Head = nil
	ll.size = 0
}

// ToSlice converts the linked list to a slice
// Complexity: O(n)
func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.size)
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Display prints the linked list
// Complexity: O(n)
func (ll *LinkedList[T]) Display() {
	if ll.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := ll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Reverse reverses the linked list in place
// Complexity: O(n)
func (ll *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := ll.Head
	var next *Node[T]

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}
