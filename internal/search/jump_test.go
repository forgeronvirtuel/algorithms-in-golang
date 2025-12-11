package search

import (
	"testing"
)

func TestJumpSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"found at beginning", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{"found at middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{"found at end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{"not found smaller", []int{1, 2, 3, 4, 5}, 0, -1},
		{"not found larger", []int{1, 2, 3, 4, 5}, 10, -1},
		{"not found between", []int{1, 3, 5, 7, 9}, 4, -1},
		{"large sorted array", []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50}, 25, 5},
		{"two elements found first", []int{1, 2}, 1, 0},
		{"two elements found second", []int{1, 2}, 2, 1},
		{"two elements not found", []int{1, 2}, 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JumpSearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("JumpSearch(%v, %d) = %d; want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkJumpSearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i * 2
	}
	target := arr[size/2]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JumpSearch(arr, target)
	}
}
