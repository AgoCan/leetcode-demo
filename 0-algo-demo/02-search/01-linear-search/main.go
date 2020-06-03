package main

import "fmt"

// 顺序查找 O(n)

func linearSearch(arrayExample []int, searchNum int) int {
	for i, v := range arrayExample {
		if v == searchNum {
			return i
		}
	}
	return -1
}

func main() {
	var arrayExample = []int{3, 4, 5, 2, 31, 23, 4}
	var searchNum = int(23)
	index := linearSearch(arrayExample, searchNum)
	fmt.Println(index)
}
