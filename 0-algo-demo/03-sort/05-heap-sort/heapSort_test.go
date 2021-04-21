package heapSortDemo

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{1, 3, 8, 4, 5, 6, 7, 9, 2, 0}
	heapSort(nums)
	fmt.Println(nums)
}
