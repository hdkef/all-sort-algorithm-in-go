package bubblesort

func BubbleSortDesc(arr []int) {
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] > arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}
}

func BubbleSortAsc(arr []int) {
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}
}
