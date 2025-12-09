package search

import "cmp"

// BinarySearch performs binary search on a sorted array
// Returns the index of target if found, otherwise -1
// Complexity: O(log n) time, O(1) space
func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs binary search recursively
func BinarySearchRecursive[T cmp.Ordered](arr []T, target T, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}

	return BinarySearchRecursive(arr, target, left, mid-1)
}
