package longestPalindromeDemo

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。

示例 2：
	输入: "cbbd"
	输出: "bb"

*/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return string(s[0])
	} else if len(s) == 0 {
		return ""
	}
	return ""
}
func expand(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		fmt.Println(left, right)
	}
	return left + 1, right - 1
}

