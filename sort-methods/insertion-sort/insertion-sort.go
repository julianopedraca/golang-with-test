package insertionsort

func InsertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		// insert A[j] into the sorted sequence A[1...j-1
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
	return arr
}
