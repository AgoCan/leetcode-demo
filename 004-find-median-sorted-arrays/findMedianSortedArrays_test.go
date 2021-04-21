package findMedianSortedArraysDemo

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	m := []int{1, 2, 3, 4}
	n := []int{2, 3}
	a := findMedianSortedArrays(m, n)
	fmt.Println(a)
}
