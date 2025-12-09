package search
package search

import (
	"math/rand"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"found at beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"found at middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"found at end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"not found", []int{1, 2, 3, 4, 5}, 10, -1},
		{"with duplicates", []int{1, 2, 3, 2, 5}, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearSearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LinearSearch(%v, %d) = %d; want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	size := 1000
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size)
	}
	target := arr[size/2]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(arr, target)
	}
}
