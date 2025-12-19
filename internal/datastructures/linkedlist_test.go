package datastructures

import (
	"testing"
)

// TestNewLinkedList tests the creation of a new linked list
func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	if ll == nil {
		t.Error("NewLinkedList should not return nil")
	}
	if ll.Head != nil {
		t.Error("New list should have nil Head")
	}
	if ll.Size() != 0 {
		t.Errorf("New list size should be 0, got %d", ll.Size())
	}
	if !ll.IsEmpty() {
		t.Error("New list should be empty")
	}
}

// TestInsertAtBeginning tests inserting elements at the beginning
func TestInsertAtBeginning(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.InsertAtBeginning(1)
	if ll.Head == nil {
		t.Fatal("Head should not be nil after insertion")
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.Value)
	}
	if ll.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ll.Size())
	}

	ll.InsertAtBeginning(2)
	if ll.Head.Value != 2 {
		t.Errorf("Expected head value 2, got %d", ll.Head.Value)
	}
	if ll.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ll.Size())
	}

	ll.InsertAtBeginning(3)
	slice := ll.ToSlice()
	expected := []int{3, 2, 1}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}
}

// TestInsertAtEnd tests inserting elements at the end
func TestInsertAtEnd(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.InsertAtEnd(1)
	if ll.Head == nil {
		t.Fatal("Head should not be nil after insertion")
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.Value)
	}

	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	slice := ll.ToSlice()
	expected := []int{1, 2, 3}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}
}

// TestInsertAtPosition tests inserting elements at specific positions
func TestInsertAtPosition(t *testing.T) {
	ll := NewLinkedList[int]()

	// Insert at position 0 (empty list)
	err := ll.InsertAtPosition(1, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.Value)
	}

	// Insert at beginning
	err = ll.InsertAtPosition(0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Insert at end
	err = ll.InsertAtPosition(3, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Insert in middle
	err = ll.InsertAtPosition(2, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	slice := ll.ToSlice()
	expected := []int{0, 1, 2, 3}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Test invalid positions
	err = ll.InsertAtPosition(5, -1)
	if err == nil {
		t.Error("Expected error for negative position")
	}

	err = ll.InsertAtPosition(5, 10)
	if err == nil {
		t.Error("Expected error for position out of bounds")
	}
}

// TestDelete tests deleting elements by value
func TestDelete(t *testing.T) {
	ll := NewLinkedList[int]()

	// Delete from empty list
	if ll.Delete(1) {
		t.Error("Delete should return false for empty list")
	}

	// Add elements
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(2)

	// Delete head
	if !ll.Delete(1) {
		t.Error("Delete should return true")
	}
	slice := ll.ToSlice()
	expected := []int{2, 3, 2}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Delete middle element (first occurrence)
	if !ll.Delete(2) {
		t.Error("Delete should return true")
	}
	slice = ll.ToSlice()
	expected = []int{3, 2}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Delete non-existent element
	if ll.Delete(5) {
		t.Error("Delete should return false for non-existent element")
	}

	// Delete remaining elements
	ll.Delete(3)
	ll.Delete(2)
	if !ll.IsEmpty() {
		t.Error("List should be empty")
	}
}

// TestDeleteAtPosition tests deleting elements at specific positions
func TestDeleteAtPosition(t *testing.T) {
	ll := NewLinkedList[int]()

	// Test invalid position on empty list
	err := ll.DeleteAtPosition(0)
	if err == nil {
		t.Error("Expected error for empty list")
	}

	// Add elements
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)

	// Delete at beginning
	err = ll.DeleteAtPosition(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	slice := ll.ToSlice()
	expected := []int{2, 3, 4}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Delete in middle
	err = ll.DeleteAtPosition(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	slice = ll.ToSlice()
	expected = []int{2, 4}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Test invalid positions
	err = ll.DeleteAtPosition(-1)
	if err == nil {
		t.Error("Expected error for negative position")
	}

	err = ll.DeleteAtPosition(10)
	if err == nil {
		t.Error("Expected error for position out of bounds")
	}
}

// TestSearch tests searching for elements
func TestSearch(t *testing.T) {
	ll := NewLinkedList[int]()

	// Search in empty list
	if ll.Search(1) {
		t.Error("Search should return false for empty list")
	}

	// Add elements
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)

	// Search for existing elements
	if !ll.Search(1) {
		t.Error("Search should find element 1")
	}
	if !ll.Search(2) {
		t.Error("Search should find element 2")
	}
	if !ll.Search(3) {
		t.Error("Search should find element 3")
	}

	// Search for non-existent element
	if ll.Search(5) {
		t.Error("Search should not find element 5")
	}
}

// TestGet tests getting elements at specific positions
func TestGet(t *testing.T) {
	ll := NewLinkedList[int]()

	// Get from empty list
	_, err := ll.Get(0)
	if err == nil {
		t.Error("Expected error for empty list")
	}

	// Add elements
	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(30)

	// Get valid positions
	val, err := ll.Get(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, err = ll.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val, err = ll.Get(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	// Get invalid positions
	_, err = ll.Get(-1)
	if err == nil {
		t.Error("Expected error for negative position")
	}

	_, err = ll.Get(10)
	if err == nil {
		t.Error("Expected error for position out of bounds")
	}
}

// TestSize tests the Size method
func TestSize(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ll.Size())
	}

	ll.InsertAtEnd(1)
	if ll.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ll.Size())
	}

	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	ll.Delete(2)
	if ll.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ll.Size())
	}

	ll.Clear()
	if ll.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", ll.Size())
	}
}

// TestIsEmpty tests the IsEmpty method
func TestIsEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	if !ll.IsEmpty() {
		t.Error("New list should be empty")
	}

	ll.InsertAtEnd(1)
	if ll.IsEmpty() {
		t.Error("List should not be empty after insertion")
	}

	ll.Delete(1)
	if !ll.IsEmpty() {
		t.Error("List should be empty after deleting all elements")
	}
}

// TestClear tests the Clear method
func TestClear(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)

	ll.Clear()

	if !ll.IsEmpty() {
		t.Error("List should be empty after Clear")
	}
	if ll.Size() != 0 {
		t.Errorf("Size should be 0 after Clear, got %d", ll.Size())
	}
	if ll.Head != nil {
		t.Error("Head should be nil after Clear")
	}
}

// TestToSlice tests the ToSlice method
func TestToSlice(t *testing.T) {
	ll := NewLinkedList[int]()

	// Empty list
	slice := ll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %v", slice)
	}

	// List with elements
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)

	slice = ll.ToSlice()
	expected := []int{1, 2, 3}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}
}

// TestReverse tests the Reverse method
func TestReverse(t *testing.T) {
	ll := NewLinkedList[int]()

	// Reverse empty list
	ll.Reverse()
	if !ll.IsEmpty() {
		t.Error("Reversed empty list should still be empty")
	}

	// Reverse single element
	ll.InsertAtEnd(1)
	ll.Reverse()
	slice := ll.ToSlice()
	expected := []int{1}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Reverse multiple elements
	ll.Clear()
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)

	ll.Reverse()
	slice = ll.ToSlice()
	expected = []int{4, 3, 2, 1}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	// Reverse again should give original order
	ll.Reverse()
	slice = ll.ToSlice()
	expected = []int{1, 2, 3, 4}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}
}

// TestLinkedListWithStrings tests the linked list with string type
func TestLinkedListWithStrings(t *testing.T) {
	ll := NewLinkedList[string]()

	ll.InsertAtEnd("hello")
	ll.InsertAtEnd("world")
	ll.InsertAtBeginning("start")

	slice := ll.ToSlice()
	expected := []string{"start", "hello", "world"}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	if !ll.Search("hello") {
		t.Error("Should find 'hello'")
	}

	if ll.Search("notfound") {
		t.Error("Should not find 'notfound'")
	}

	ll.Delete("hello")
	if ll.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ll.Size())
	}
}

// TestLinkedListComplexOperations tests a sequence of complex operations
func TestLinkedListComplexOperations(t *testing.T) {
	ll := NewLinkedList[int]()

	// Build a list: [1, 2, 3, 4, 5]
	for i := 1; i <= 5; i++ {
		ll.InsertAtEnd(i)
	}

	// Insert 0 at beginning: [0, 1, 2, 3, 4, 5]
	ll.InsertAtBeginning(0)

	// Insert 10 at position 3: [0, 1, 2, 10, 3, 4, 5]
	ll.InsertAtPosition(10, 3)

	// Delete element 3: [0, 1, 2, 10, 4, 5]
	ll.Delete(3)

	// Delete at position 1: [0, 2, 10, 4, 5]
	ll.DeleteAtPosition(1)

	slice := ll.ToSlice()
	expected := []int{0, 2, 10, 4, 5}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}

	if ll.Size() != 5 {
		t.Errorf("Expected size 5, got %d", ll.Size())
	}

	// Reverse: [5, 4, 10, 2, 0]
	ll.Reverse()
	slice = ll.ToSlice()
	expected = []int{5, 4, 10, 2, 0}
	if !sliceEqual(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}
}

// Helper function to compare slices
func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
