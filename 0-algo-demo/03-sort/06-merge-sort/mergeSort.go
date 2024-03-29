package mergeSortDemo

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
	// 左右依次做比较，因为递归后最小单位是2个元素的数组，就能保证每一次比较的数组都是有序的
	for i <= mid && j <= high {
		if nums[i] < nums[j] {
			numsTmp = append(numsTmp, nums[i])
			i++
		} else {
			numsTmp = append(numsTmp, nums[j])
			j++
		}
	}
	// 左侧的还有数据，就加载临时数据的后面
	for ; i <= mid; i++ {
		numsTmp = append(numsTmp, nums[i])
	}
	// 右侧的还有数据，就加载临时数据的后面
	for ; j <= high; j++ {
		numsTmp = append(numsTmp, nums[j])
	}
	// 排序好之后直接赋值原来的切片
	for k := low; k < len(numsTmp); k++ {
		nums[k] = numsTmp[k]
	}
}

func mergeSort(nums []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		// 左侧比较
		mergeSort(nums, low, mid)
		// 右侧比较
		mergeSort(nums, mid+1, high)

		merge(nums, low, mid, high)
	}

}
