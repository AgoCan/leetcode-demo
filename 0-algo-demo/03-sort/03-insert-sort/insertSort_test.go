package insertSortDemo

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	a := []int{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := insertSort(a)
	fmt.Println(b)
}
