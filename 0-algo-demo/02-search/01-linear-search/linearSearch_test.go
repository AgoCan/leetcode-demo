package linearSearchDemo

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	var arrayExample = []int{3, 4, 5, 2, 31, 23, 4}
	var searchNum = int(23)
	index := linearSearch(arrayExample, searchNum)
	fmt.Println(index)
}
