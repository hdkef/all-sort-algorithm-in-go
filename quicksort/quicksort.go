package quicksort

import "fmt"

func partitionAsc(arr []int, left int, right int, pivot int) int {
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func QuickSortAsc(arr []int, left int, right int) {
	if left >= right {
		return
	}

	pivot := arr[(left+right)/2]
	idx := partitionAsc(arr, left, right, pivot)
	QuickSortAsc(arr, left, idx-1)
	QuickSortAsc(arr, idx, right)
}

func partitionDesc(arr []int, left int, right int, pivot int) int {
	for left <= right {
		for arr[left] > pivot {
			left++
		}

		for arr[right] < pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	fmt.Println(arr)
	return left
}

func QuickSortDesc(arr []int, left int, right int) {
	if left >= right {
		return
	}

	pivot := arr[(left+right)/2]
	idx := partitionDesc(arr, left, right, pivot)
	QuickSortDesc(arr, left, idx-1)
	QuickSortDesc(arr, idx, right)
}
