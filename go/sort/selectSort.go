package sort

func SelectSortDesc(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[index] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

func SelectSortAsc(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}
