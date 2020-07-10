package main

import "fmt"

// 归并, 此代码有问题，
/*
[1,3,5,7,9]
[2,4,6,8]
两个列表合并
[1,3,5,7,9,2,4,6,8]
low是左边列表的下标
mid是左边列表的最右下标
high是整个列表的最右下标
[1,2,3,4,5,6,7,8,9]
*/
func merge(nums []int, low, mid, high int) {
	i := low
	j := mid + 1
	var numsTmp []int
	for i <= mid && j <= high {

		if nums[i] < nums[j] {
			numsTmp = append(numsTmp, nums[i])
			i++
		} else {
			numsTmp = append(numsTmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		numsTmp = append(numsTmp, nums[i])

	}
	for ; j <= high; j++ {
		numsTmp = append(numsTmp, nums[j])

	}
	for k := low; k < len(numsTmp); k++ {
		nums[k] = numsTmp[k]
	}
}

func mergeSort(nums []int, low, high int) {
	if low < high {

		mid := (low + high) / 2
		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)
		fmt.Println(nums[low : high+1])
		merge(nums, low, mid, high)

	}

}

func main() {
	nums := []int{1, 3, 5, 7, 2, 3, 2, 3, 2, 4, 6, 8, 9, 10}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
