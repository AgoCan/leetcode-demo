package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
func twoSum(nums []int, target int) (ret []int) {
	var res map[int]int
	res = make(map[int]int, len(nums))
	ret = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		desired := target - nums[i]
		fmt.Println(desired, i)
		_, ok := res[desired]
		if !ok {
			res[nums[i]] = i
		} else {
			ret = []int{res[desired], i}
			return ret
		}
	}
	return ret
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)

}
