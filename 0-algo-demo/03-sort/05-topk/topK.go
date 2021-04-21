package topKDemo

// 小根堆
func sift(nums []int, low, high int) []int {
	i := low         // 堆的根节点
	j := 2*i + 1     // j 开始是左孩子
	tmp := nums[low] // 把堆的顶点存起来
	for {
		if j > high { // 只要j的位置有数
			nums[i] = tmp
			break
		}
		if j+1 <= high && nums[j+1] < nums[j] { // 和大根堆的区别一小于号，如果有，且右孩子比较小，j指向右孩子
			j = j + 1
		}
		if nums[j] < tmp { // 和大根堆的区别二， 这里也是小于号
			nums[i] = nums[j]
			i = j
			j = 2*i + 1
		} else {
			nums[i] = tmp
			break
		}
	}
	return nums
}

func topK(nums []int, k int) []int {
	heap := nums[0:k]
	// 建堆
	for i := (k - 2) / 2; i > -1; i-- {
		sift(heap, i, k-1)
	}
	// 遍历所有数
	for i := k; i < len(nums)-1; i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			sift(heap, 0, k-1)
		}
	}
	// 出数
	for i := k - 1; i > -1; i-- {
		// i 一直指向堆的最后一个位置
		heap[0], heap[i] = heap[i], heap[0]
		sift(heap, 0, i-1)
	}
	return heap
}
