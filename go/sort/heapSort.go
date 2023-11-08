package sort

func SmallHeapSort(arr []int32) []int32 {
	arr = createSmallHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		tArr := arr[:i]
		downSmallAdjust(0, &tArr)
	}
	return arr
}

func BigHeapSort(arr []int32) []int32 {
	arr = createBigHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		tArr := arr[:i]
		downBigAdjust(0, &tArr)
	}
	return arr
}

func createBigHeap(arr []int32) []int32 {
	for i := len(arr) / 2; i >= 0; i-- {
		downBigAdjust(i, &arr)
	}
	return arr
}

func createSmallHeap(arr []int32) []int32 {
	for i := len(arr) / 2; i >= 0; i-- {
		downSmallAdjust(i, &arr)
	}
	return arr
}

func downSmallAdjust(index int, array *[]int32) {
	arr := *array
	if index >= len(arr)/2 {
		return
	}

	var (
		left  = index*2 + 1
		right = index*2 + 2
	)

	if right < len(arr) && arr[right] < arr[left] {
		if arr[right] < arr[index] {
			arr[right], arr[index] = arr[index], arr[right]
			downSmallAdjust(right, array)
		}
		return
	}

	if arr[left] < arr[index] {
		arr[left], arr[index] = arr[index], arr[left]
		downSmallAdjust(left, array)
	}
}

func downBigAdjust(index int, array *[]int32) {
	arr := *array
	if index >= len(arr)/2 {
		return
	}
	var (
		left  = index*2 + 1
		right = index*2 + 2
	)
	if right < len(arr) && arr[left] < arr[right] {
		if arr[right] > arr[index] {
			arr[right], arr[index] = arr[index], arr[right]
			downBigAdjust(index*2+2, array)
		}
		return
	}
	if arr[left] > arr[index] {
		arr[left], arr[index] = arr[index], arr[left]
		downBigAdjust(index*2+1, array)
	}
}
