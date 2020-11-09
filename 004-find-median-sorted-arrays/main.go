package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

示例 1:
	输入：nums1 = [1,3], nums2 = [2]
	输出：2.00000
	解释：合并数组 = [1,2,3] ，中位数 2

示例 2:
	输入：nums1 = [1,2], nums2 = [3,4]
	输出：2.50000
	解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：
	输入：nums1 = [0,0], nums2 = [0,0]
	输出：0.00000
示例 4：
	输入：nums1 = [], nums2 = [1]
	输出：1.00000
示例 5：
	输入：nums1 = [2], nums2 = []
	输出：2.00000
提示：
	nums1.length == m
	nums2.length == n
	0 <= m <= 1000
	0 <= n <= 1000
	1 <= m + n <= 2000
	-10^6 <= nums1[i], nums2[i] <= 10^6
*/

// 二分查找

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	// 其中 midIndex + 1 是保证左边比右边的数字多一位
	// 两个数组长度和 奇数
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	// 两个数组长度和 偶数
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	// (左 + 右)/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			// 如果数组 1 的长度为0， 则直接计算数组2的中位数即可
			// [1,2,3]  mid = 1,  0+(1+1)-1
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			// 两个数组均只有一个数字，返回左边的数字
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		// k = 3
		// a = [1,2,3,4]  b = [2,3]
		// half = 1  index1=0   1 < 4  new1 = 0
		// half = 1  index2=0   1 < 2  new2 = 0
		// a[0] < b[0]
		// k = 3 - 0 + 1
		// k = 4
		// half = 1  index1=1  2 < 4  new1 = 1
		// half = 1  index2=0  1 < 2  new2 = 0
		// a[1] = b[1]
		// k = 1 (因为上面计算了 k==1的原因，所以此处不执行)
		// 第二轮
		// k = 4
		// half = 2  index1=0   2 < 4  new1 = 1
		// half = 2  index2=0   2 = 2  new2 = 1
		// a[1] < b[1]
		// k = 4 - 0 + 1
		
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		fmt.Println(half, newIndex1, newIndex2)
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			// 由于 左 < 右
			// 向右移动一位
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	m := []int{1, 2, 3, 4}
	n := []int{2, 3}
	a := findMedianSortedArrays(m, n)
	fmt.Println(a)
}
