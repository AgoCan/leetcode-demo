package main

import "fmt"

// 选择排序 O(n^2)

func selectionSort(nums []int32) []int32 {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
		fmt.Println(nums)
	}
	return nums
}

func main() {
	a := []int32{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := selectionSort(a)
	fmt.Println(b)
}
