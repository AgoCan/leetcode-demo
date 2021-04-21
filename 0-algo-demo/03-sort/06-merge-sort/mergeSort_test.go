package mergeSortDemo

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{1, 3, 5, 7, 2, 3, 2, 3, 2, 4, 6, 8, 9, 10}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
