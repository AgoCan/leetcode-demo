package twoSum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 1, 8}
	target := 3
	res := Sum(nums, target)
	fmt.Println(res)
}
