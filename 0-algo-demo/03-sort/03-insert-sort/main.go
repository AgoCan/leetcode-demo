package main

import "fmt"

// 插入排序 O(n^2)

func insertSort(nums []int) []int {
	length := len(nums)
	var j int
	var temp int
	for i := 1; i < length; i++ {
		temp = nums[i]
		j = i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
	return nums
}

func main() {
	a := []int{1, 2, 5, 2, 4, 6, 8, 8, 55, 4, 33, 2}
	b := insertSort(a)
	fmt.Println(b)
}
