package insertSortDemo

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
