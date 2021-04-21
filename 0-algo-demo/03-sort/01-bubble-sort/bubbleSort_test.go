package bubbleSortDemo

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := bubbleSort(a)
	fmt.Println(b)
}
