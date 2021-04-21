package linearSearchDemo

// 顺序查找 O(n)

func linearSearch(arrayExample []int, searchNum int) int {
	for i, v := range arrayExample {
		if v == searchNum {
			return i
		}
	}
	return -1
}
