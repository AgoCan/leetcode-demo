package hanoiDemo

import "fmt"

// 递归

func hanoi(n, a, b, c int) {

	if n > 0 {
		hanoi(n-1, a, c, b)
		fmt.Printf("moving from %v to %v\n", a, c)
		hanoi(n-1, b, a, c)
	}
}
