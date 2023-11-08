package sort

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"sort"
	"testing"
)

const (
	ARR_LENGTH = 1000
)

var (
	arr1      = make([]int32, ARR_LENGTH)
	arr1_asc  = make([]int32, ARR_LENGTH)
	arr1_desc = make([]int32, ARR_LENGTH)
)

func init() {
	genRandomInt32Slice(&arr1)
	copy(arr1_asc, arr1)
	copy(arr1_desc, arr1)
	sort.Slice(arr1_asc, func(i, j int) bool {
		return arr1_asc[i] < arr1_asc[j]
	})
	sort.Slice(arr1_desc, func(i, j int) bool {
		return arr1_desc[i] > arr1_desc[j]
	})
}

func genRandomInt32Slice(res *[]int32) {
	for i := 0; i < len(*res); i++ {
		(*res)[i] = rand.Int31n(math.MaxInt32)
	}
}

func TestBubbleSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, BubbleSortAsc(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, BubbleSortDesc(desc))
}

func TestSelectSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, SelectSortAsc(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, SelectSortDesc(desc))
}

func TestQuickSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, QuickSortAsc(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, QuickSortDesc(desc))
}

func TestMergeSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, MergeSortAsc(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, MergeSortDesc(desc))
}

func TestInsertionSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, InsertionSortAsc(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, InsertionSortDesc(desc))
}

func TestHeapSort(t *testing.T) {
	var (
		asc  = make([]int32, ARR_LENGTH)
		desc = make([]int32, ARR_LENGTH)
	)
	copy(asc, arr1)
	assert.Equal(t, arr1_asc, BigHeapSort(asc))

	copy(desc, arr1)
	assert.Equal(t, arr1_desc, SmallHeapSort(desc))
}
