package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 3, 2, 1, 2}, []int{1, 2, 2, 3, 3}},
	}

	for _, tt := range tests {
		arr := make([]int, len(tt.input))
		copy(arr, tt.input)
		Quicksort(arr, 0, len(arr)-1)
		if !reflect.DeepEqual(arr, tt.expected) {
			t.Errorf("Quicksort(%v) = %v; want %v", tt.input, arr, tt.expected)
		}
	}
}

func BenchmarkQuicksort(b *testing.B) {
	size := 1000
	for i := 0; i < b.N; i++ {
		arr := make([]int, size)
		for j := range arr {
			arr[j] = rand.Intn(size)
		}
		Quicksort(arr, 0, len(arr)-1)
	}
}
