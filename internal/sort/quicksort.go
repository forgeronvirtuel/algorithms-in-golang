package sort

func Quicksort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		Quicksort(arr, low, p-1)
		Quicksort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose last element as pivot
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // return pivot index
}
