package datastructures

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue single element
	q.Enqueue(10)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}

	// Enqueue multiple elements
	q.Enqueue(20)
	q.Enqueue(30)
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	// Dequeue from empty queue
	val, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue from empty queue should return false")
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}

	// Enqueue and dequeue single element
	q.Enqueue(10)
	val, ok = q.Dequeue()
	if !ok {
		t.Error("Dequeue should return true")
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing last element")
	}

	// Test FIFO order
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, _ = q.Dequeue()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	val, _ = q.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	val, _ = q.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue[string]()

	// Peek on empty queue
	val, ok := q.Peek()
	if ok {
		t.Error("Peek on empty queue should return false")
	}
	if val != "" {
		t.Errorf("Expected empty string, got %s", val)
	}

	// Peek on non-empty queue
	q.Enqueue("first")
	q.Enqueue("second")

	val, ok = q.Peek()
	if !ok {
		t.Error("Peek should return true on non-empty queue")
	}
	if val != "first" {
		t.Errorf("Expected 'first', got %s", val)
	}

	// Verify size didn't change after peek
	if q.Size() != 2 {
		t.Errorf("Expected size 2 after peek, got %d", q.Size())
	}
}

func TestQueueSize(t *testing.T) {
	q := NewQueue[int]()

	// Empty queue
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}

	// Add elements
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
		if q.Size() != i {
			t.Errorf("Expected size %d, got %d", i, q.Size())
		}
	}

	// Remove elements
	for i := 5; i > 0; i-- {
		q.Dequeue()
		if q.Size() != i-1 {
			t.Errorf("Expected size %d, got %d", i-1, q.Size())
		}
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue[int]()

	// Initially empty
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	// Not empty after enqueue
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}

	// Empty after dequeue
	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func TestQueueWithDifferentTypes(t *testing.T) {
	// Test with strings
	qString := NewQueue[string]()
	qString.Enqueue("hello")
	qString.Enqueue("world")
	val, _ := qString.Dequeue()
	if val != "hello" {
		t.Errorf("Expected 'hello', got %s", val)
	}

	// Test with float64
	qFloat := NewQueue[float64]()
	qFloat.Enqueue(3.14)
	qFloat.Enqueue(2.71)
	valFloat, _ := qFloat.Dequeue()
	if valFloat != 3.14 {
		t.Errorf("Expected 3.14, got %f", valFloat)
	}

	// Test with custom struct
	type Person struct {
		Name string
		Age  int
	}
	qPerson := NewQueue[Person]()
	qPerson.Enqueue(Person{Name: "Alice", Age: 30})
	qPerson.Enqueue(Person{Name: "Bob", Age: 25})
	person, _ := qPerson.Dequeue()
	if person.Name != "Alice" || person.Age != 30 {
		t.Errorf("Expected Alice/30, got %s/%d", person.Name, person.Age)
	}
}

func TestQueueEnqueueDequeueAlternating(t *testing.T) {
	q := NewQueue[int]()

	// Alternating enqueue and dequeue operations
	q.Enqueue(1)
	val, _ := q.Dequeue()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	q.Enqueue(2)
	q.Enqueue(3)
	val, _ = q.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	q.Enqueue(4)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	val, _ = q.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	val, _ = q.Dequeue()
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestQueueLargeNumberOfElements(t *testing.T) {
	q := NewQueue[int]()
	n := 1000

	// Enqueue n elements
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}

	if q.Size() != n {
		t.Errorf("Expected size %d, got %d", n, q.Size())
	}

	// Dequeue and verify order
	for i := 0; i < n; i++ {
		val, ok := q.Dequeue()
		if !ok {
			t.Errorf("Dequeue failed at iteration %d", i)
		}
		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}
