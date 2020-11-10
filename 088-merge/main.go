package main

import "fmt"

/*
合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明：
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例：
	输入：
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	输出：[1,2,2,3,5,6]
提示：
	-10^9 <= nums1[i], nums2[i] <= 10^9
	nums1.length == m + n
	nums2.length == n
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])
	p1 := 0
	p2 := 0
	index := 0
	for p1 < m && p2 < n {
		if nums1Copy[p1] < nums2[p2] {

			nums1[index] = nums1Copy[p1]
			p1++
			index++
		} else {
			nums1[index] = nums2[p2]
			p2++
			index++
		}
		fmt.Println(nums1)
	}
	if p1 < m {
		for ; p1 < m; p1++ {
			nums1[index] = nums1Copy[p1]
			index++
		}
	} else {
		for ; p2 < n; p2++ {
			nums1[index] = nums2[p2]
			index++
		}
	}
	fmt.Println(nums1)
}

func main() {
	a := []int{4, 5, 6, 0, 0, 0}
	b := []int{1, 2, 3}

	merge(a, 3, b, 3)
}
