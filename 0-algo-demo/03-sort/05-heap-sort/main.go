package main

import "fmt"

/*
最大堆：每个父节点的都大于孩子节点。
最小堆：每个父节点的都小于孩子节点。

							9
			8								7
	4				5				6				3
1		2		0

满二叉树， 最下面的数字填满，完全二叉树，最下面的叶子节点，从左到右填满
				    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
列表展示上面完全二叉树  [9, 8, 7, 4, 5, 6, 3, 1, 2, 0]

父节点找（左）子节点的下标 i = 2i+1
父节点找（右）子节点的下标 i = 2i+2

9的下标是0， 8的下标是1   i = 2*0 + 1 = 1
9的下标是0， 7的下标是2   i = 2*0 + 2 = 2
8的下标是1， 4的下标是3   i = 2*1 + 1 = 3
7的下标是2， 6的下标是5	  i = 2*2 + 1 = 5
以此类推，找下标

孩子找父亲
不管左右孩子： i = (i-1)/2, 左下标都是偶数，右下标都是基数。 基数 处以 2 经过int之后就变成了跟偶数一样

*/

/*
	low,  堆的根节点位置
	high, 堆的最后一个元素位置
	为什么要有high？ 因为，把调整后数组的第一个元素 放在该数组的最后一个位置上，具体根据代码进行分析
*/
// sift 向下调整
func sift(nums []int, low, high int) []int {
	i := low         // 堆的根节点
	j := 2*i + 1     // j 开始是左孩子
	tmp := nums[low] // 把堆的顶点存起来
	for {
		if j > high { // 只要j的位置有数
			nums[i] = tmp
			break
		}
		if j+1 <= high && nums[j+1] > nums[j] { // 如果有，且右孩子比较大，j指向右孩子
			j = j + 1
		}
		if nums[j] > tmp {
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

func heapSort(nums []int) {
	n := len(nums)
	// (n-1) 最后一个数字的下标
	// ((n-1)-1) / 2 等与父亲节点，把n-1想成i 即 (i-1)/2
	// 建立最大堆
	for i := (n - 2) / 2; i > -1; i-- {
		sift(nums, i, n-1)

	}
	for i := n - 1; i > -1; i-- {
		// i 一直指向堆的最后一个位置
		nums[0], nums[i] = nums[i], nums[0]
		sift(nums, 0, i-1)
	}

}

func main() {
	nums := []int{1, 3, 8, 4, 5, 6, 7, 9, 2, 0}
	heapSort(nums)
	fmt.Println(nums)
}
