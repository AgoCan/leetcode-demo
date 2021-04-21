package bubbleSortDemo

// 冒泡排序 O(n^2)

func bubbleSort(sli []int) []int {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1-i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	return (sli)
}
