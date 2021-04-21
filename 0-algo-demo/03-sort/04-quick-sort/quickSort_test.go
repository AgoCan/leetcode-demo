package quickSortDemo

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := quickSort(a)
	fmt.Println(b)
}
