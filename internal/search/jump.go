package search

import "math"

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	step := int(math.Sqrt(float64(n)))
	prev := 0

	// Step 1: Jump forward in blocks
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	// Step 2: Linear search inside the block
	for prev < min(step, n) {
		if arr[prev] == target {
			return prev
		}
		if arr[prev] > target {
			return -1
		}
		prev++
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
