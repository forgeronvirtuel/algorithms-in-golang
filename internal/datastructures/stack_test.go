package datastructures

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	if stack == nil {
		t.Error("Expected new stack to be created, got nil")
	}
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty")
	}
}

func TestStackPushPopPeek(t *testing.T) {
	stack := NewStack[int]()

	// Test Push
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Errorf("Expected size 3 after pushes, got %d", stack.Size())
	}

	// Test Peek
	top, ok := stack.Peek()
	if !ok || top != 30 {
		t.Errorf("Expected top to be 30, got %d", top)
	}

	// Test Pop
	value, ok := stack.Pop()
	if !ok || value != 30 {
		t.Errorf("Expected popped value to be 30, got %d", value)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}

	// Pop remaining elements
	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after popping all elements")
	}

	// Test Pop on empty stack
	_, ok = stack.Pop()
	if ok {
		t.Error("Expected Pop on empty stack to return false")
	}

	// Test Peek on empty stack
	_, ok = stack.Peek()
	if ok {
		t.Error("Expected Peek on empty stack to return false")
	}
}

func TestStackWithDifferentTypes(t *testing.T) {
	// Test with string type
	stringStack := NewStack[string]()
	stringStack.Push("hello")
	stringStack.Push("world")

	top, ok := stringStack.Peek()
	if !ok || top != "world" {
		t.Errorf("Expected top to be 'world', got '%s'", top)
	}

	value, ok := stringStack.Pop()
	if !ok || value != "world" {
		t.Errorf("Expected popped value to be 'world', got '%s'", value)
	}

	// Test with custom struct type
	type Point struct {
		X, Y int
	}

	pointStack := NewStack[Point]()
	pointStack.Push(Point{1, 2})
	pointStack.Push(Point{3, 4})

	topPoint, ok := pointStack.Peek()
	if !ok || topPoint != (Point{3, 4}) {
		t.Errorf("Expected top to be Point{3, 4}, got %+v", topPoint)
	}

	valuePoint, ok := pointStack.Pop()
	if !ok || valuePoint != (Point{3, 4}) {
		t.Errorf("Expected popped value to be Point{3, 4}, got %+v", valuePoint)
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty after push")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after popping the only element")
	}
}

func TestStackSize(t *testing.T) {
	stack := NewStack[int]()

	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}

	stack.Push(1)
	stack.Push(2)
	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pushes, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 0 {
		t.Errorf("Expected size 0 after popping all elements, got %d", stack.Size())
	}
}
