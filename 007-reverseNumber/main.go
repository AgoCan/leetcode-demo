package main

import "fmt"

/*给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

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

func main() {
	fmt.Println(reverse(1))

}
