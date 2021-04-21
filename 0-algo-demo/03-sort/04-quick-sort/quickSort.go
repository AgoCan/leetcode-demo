package quickSortDemo

// 快速排序 O(nlogn)

// 拿出第一个元素，然后 保证左边的数字都比他小，右边的数字都比他大，然后归位（赋值）
func partition(nums []int, left, right int) (mid int) {
	temp := nums[left]
	for left < right {
		// 左边有一个空格，所以在右边找
		for nums[right] >= temp && left < right {
			right--
		}
		nums[left] = nums[right]

		for nums[left] < temp && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = temp // 归位

	mid = left
	return mid
}

func rQuickSort(nums []int, left, right int) []int {
	var mid int
	if left < right {
		mid = partition(nums, left, right)
		rQuickSort(nums, left, mid-1)
		rQuickSort(nums, mid+1, right)
	}
	return nums
}
func quickSort(nums []int) []int {
	rQuickSort(nums, 0, len(nums)-1)
	return nums
}
