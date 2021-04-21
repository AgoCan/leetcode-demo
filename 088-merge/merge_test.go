package mergeDemo

import "testing"

func TestMerge(t *testing.T) {
	a := []int{4, 5, 6, 0, 0, 0}
	b := []int{1, 2, 3}

	merge(a, 3, b, 3)
}
