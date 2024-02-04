package mergesort

func Mergesort(array []int) []int {
	length := len(array)

	if length <= 1 {
		return array
	}

	middle := length / 2
	leftArray := make([]int, middle)
	rightArray := make([]int, length-middle)

	j := 0
	for i := 0; i < length; i++ {
		if i < middle {
			leftArray[i] = array[i]
		} else {
			rightArray[j] = array[i]
			j++
		}
	}

	Mergesort(leftArray)
	Mergesort(rightArray)
	merge(leftArray, rightArray, array)

	return array
}

func merge(leftArray []int, rightArray []int, array []int) {
	leftSize := len(array) / 2
	rightSize := len(array) - leftSize

	i, l, r := 0, 0, 0 // indieces

	for l < leftSize && r < rightSize {
		if leftArray[l] < rightArray[r] {
			array[i] = leftArray[l]
			i++
			l++
		} else {
			array[i] = rightArray[r]
			i++
			r++
		}
	}
	for l < leftSize {
		array[i] = leftArray[l]
		i++
		l++
	}
	for r < rightSize {
		array[i] = rightArray[r]
		i++
		r++
	}
}
