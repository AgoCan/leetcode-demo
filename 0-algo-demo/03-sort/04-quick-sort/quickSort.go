package quickSortDemo

// 快速排序 O(nlogn)

func partition(nums []int, left, right int) (mid int) {
	// 拿左数来做 中枢数，这个数字最终最放在属于它的位置，左边都比它小，右边都比它大
	temp := nums[left]

	for left < right {
		// 左边有一个空格，所以在右边找
		// 右边的数据，大于就继续放在右边，如果小于，就去放在
		for nums[right] >= temp && left < right {
			right--
		}
		// 把右边的数字放在左下标，这样temp的原本的数据就被覆盖了
		nums[left] = nums[right]

		// 右边的数据移动到左边，所以需要在左边在取一个数据给放回来
		for nums[left] < temp && left < right {
			left++
		}
		nums[right] = nums[left]
		// 左边一直加，右边一直减，一旦碰面就是 中枢数的下标
	}
	nums[left] = temp // 归位

	mid = left
	return mid
}

func rQuickSort(nums []int, left, right int) []int {
	var mid int
	if left < right {
		// mid不是真正中心节点，而是中枢的数据的下标，这个下标的左边都比它小，右边都比它大
		mid = partition(nums, left, right)
		// 然后 按照mid进行切分，左右两边无序的数据重新排序
		rQuickSort(nums, left, mid-1)
		rQuickSort(nums, mid+1, right)
	}
	return nums
}

func quickSort(nums []int) []int {
	rQuickSort(nums, 0, len(nums)-1)
	return nums
}
