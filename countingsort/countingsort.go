package countingsort

func CountingSortDesc(arr []int, k int) {
	countArr := make([]int, k+1) //to automate fill the index with 0 and give len k+1
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}
	counter := len(arr) - 1
	for i := 0; i < k+1; i++ {
		for j := 0; j < countArr[i]; j++ {
			arr[counter] = i
			counter--
		}
	}
}

func CountingSortAsc(arr []int, k int) {
	countArr := make([]int, k+1) //to automate fill the index with 0 and give len k+1
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
	}
	counter := 0
	for i := 0; i < k+1; i++ {
		for j := 0; j < countArr[i]; j++ {
			arr[counter] = i
			counter++
		}
	}
}
