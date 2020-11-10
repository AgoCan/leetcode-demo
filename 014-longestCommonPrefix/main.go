package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
	输入: ["flower","flow","flight"]
	输出: "fl"

示例 2:
	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。

*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		// 如果有空字符串，直接退出即可
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	// 公共前缀最长不会超过 最短的字符串
	length := min(len(str1), len(str2))
	index := 0
	// 每一个字符串都进行比较，不相等就说再见了
	for index < length && str1[index] == str2[index] {
		index++
	}
	// 如果为0，直接退出即可
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
