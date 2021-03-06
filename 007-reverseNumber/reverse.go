package reverseDemo

/*
简单
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

*/

func reverse(x int) int {
	y := 0
	flag := 1
	if x < 0 {
		x = -x
		flag = -flag
	}
	if x > 4294967296 {
		return 0
	}
	for {
		if x < 10 {
			cur := x % 10
			y = y*10 + cur
			break
		}
		cur := x % 10
		y = y*10 + cur
		x /= 10

	}
	if y > 2147483648 {
		return 0
	}
	if flag < 0 {
		y = -y
	}
	return y
}
