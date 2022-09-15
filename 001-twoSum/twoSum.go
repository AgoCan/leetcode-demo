package twoSum

/*
简单

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

/*
解题思路：
	有序：     [2, 3, 5,7] 得8
	左右相加，数值大于 8，右index减1， 小于，左减1

	无序：     [3, 5, 2, 7] 得8 map[int]int    map[value]=index
	map[3:0, 5:1]    3+5 =8  得到index等于0和1

	无序：     [3, 7, 2, 5] 得8 map[int]int    map[value]=index
	map[3:0, 5:1, 2:2, 5:3]    3+5 =8  得到index等于0和3
*/
func Sum(nums []int, target int) (ret []int) {

	res := make(map[int]int, len(nums))
	ret = make([]int, 2)

	for i, v := range nums {
		// i 是index  目标减去值
		desired := target - v
		// 如果目标减去下标所对应的值，就记录下来
		// 如果目标存在。说明找到了值，可以返回了
		// 例如 [4,2,1,3] 等于5， 第一轮找到了  4对0，2对1，1对2， 5-1=4在map里面
		_, ok := res[desired]
		if !ok {
			// value = index
			res[v] = i
		} else {
			ret = []int{res[desired], i}
			return ret
		}
	}
	return ret
}
