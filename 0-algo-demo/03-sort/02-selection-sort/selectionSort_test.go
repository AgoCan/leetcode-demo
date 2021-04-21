package selectionSortDemo

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int32{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := selectionSort(a)
	fmt.Println(b)
}
