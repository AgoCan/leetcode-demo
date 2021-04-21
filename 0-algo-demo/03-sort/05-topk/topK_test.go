package topKDemo

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	nums := []int{1, 3, 8, 4, 5, 6, 7, 9, 2, 0}
	fmt.Println(topK(nums, 3))
}
