package sortalgo

//go:generate go run ./cmd/generate_testdata/main.go

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, tt := range GeneratedTestData {
		arr := make([]int, len(tt.Unsorted))
		copy(arr, tt.Unsorted)
		InsertionSort(arr)
		if !reflect.DeepEqual(arr, tt.Sorted) {
			t.Errorf("InsertionSort(%v) = %v; want %v", tt.Unsorted, arr, tt.Sorted)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	size := 1000
	for b.Loop() {
		arr := make([]int, size)
		for j := range arr {
			arr[j] = rand.Intn(size)
		}
		InsertionSort(arr)
	}
}
