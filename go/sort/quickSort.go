package sort

import "math/rand"

func QuickSortDesc(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}
	var (
		index       = rand.Intn(len(arr))
		value       = arr[index]
		left, right = 0, len(arr) - 1
	)
	arr[index], arr[left] = arr[left], arr[index]
	for left < right {
		for ; arr[right] < value && left < right; right-- {
		}
		arr[left] = arr[right]
		for ; arr[left] > value && left < right; left++ {

		}
		arr[right] = arr[left]
	}

	arr[left] = value
	leftArr := QuickSortDesc(arr[:left])
	rightArr := QuickSortDesc(arr[left+1:])
	return append(append(leftArr, arr[left]), rightArr...)
}

func QuickSortAsc(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}
	var (
		length = len(arr)
		index  = rand.Intn(length - 1)
		val    = arr[index]
		left   = 0
		right  = length - 1
	)
	arr[left], arr[index] = arr[index], arr[left]
	for left < right {
		for ; arr[right] > val && left < right; right-- {
		}
		arr[left] = arr[right]
		for ; arr[left] < val && left < right; left++ {
		}
		arr[right] = arr[left]
	}
	arr[left] = val
	leftArr := QuickSortAsc(arr[:left])
	rightArr := QuickSortAsc(arr[left+1:])
	return append(append(leftArr, arr[left]), rightArr...)
}
