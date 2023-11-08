package sort

func MergeSortAsc(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}
	length := len(arr)
	left := MergeSortAsc(arr[:length/2])
	right := MergeSortAsc(arr[length/2:])
	return mergeAsc(left, right)
}

func MergeSortDesc(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}

	left := MergeSortDesc(arr[:len(arr)/2])
	right := MergeSortDesc(arr[len(arr)/2:])
	return mergeDesc(left, right)
}

func mergeDesc(left, right []int32) []int32 {
	var (
		i, j   int
		result []int32
	)
	for i < len(left) && j < len(right) {
		for ; i < len(left) && j < len(right) && left[i] > right[j]; i++ {
			result = append(result, left[i])
		}
		for ; j < len(right) && i < len(left) && right[j] > left[i]; j++ {
			result = append(result, right[j])
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	for ; j < len(right); j++ {
		result = append(result, right[j])
	}
	return result
}

func mergeAsc(left, right []int32) []int32 {
	var result []int32
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
