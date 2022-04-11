package main

//给你一个字符串 s，找到 s 中最长的回文子串。
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。

func longestPalindrome(s string) string {
	var maxS string
	for i := 0; i < len(s); i++ {
		s1, s2 := getMaxLen(s, i, i), getMaxLen(s, i, i+1)
		if len(s1) > len(maxS) {
			maxS = s1
		}
		if len(s2) > len(maxS) {
			maxS = s2
		}
	}
	return maxS
}

func getMaxLen(s string, i, j int) string {
	l, r := 0, len(s)-1
	for i >= l && j <= r {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}

	return s[i+1 : j]
}

//中心扩展：枚举每个奇偶长度的回文串的中心起点，找最长的可能性
