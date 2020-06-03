package main

// 二分查找,查找次数多，需要先排序。不然使用线性查找  O(logn)

func binarySearch(arrayExample []int, searchNum int) int {
	left := 0
	right := len(arrayExample) - 1
	for left < right {
		mid := (left + right) / 2
		if arrayExample[mid] == searchNum {
			return mid
		} else if arrayExample[mid] > searchNum {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {

}
