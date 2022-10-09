package longestCommonPrefixDemo

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	a := []string{
		"flower", "flow", "flight",
	}
	fmt.Println(longestCommonPrefix(a))
}
