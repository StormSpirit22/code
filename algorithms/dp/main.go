package main

import "fmt"

func main() {
	//fmt.Println(minDistance("sea", "eat"))
	//fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome2("cbbd"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	// dp[i][j]: word1 在索引 i 和 word2 在索引 j 时需要删除的字符个数
	// base case: 当 word1 是空或 word2 是空时，值就是另一个字符串的长度，dp[0][j] = j, dp[i][0] = i
	// func:
	// 1. if word1[i] == word2[j] 则不需要删除。dp[i][j] = dp[i-1][j-1]
	// 2. else 选择删除 word1 和 word2 中的某一个的当前字符，选最小值并+1。 dp[i][j] = min(dp[i-1][j], dp[j-1][i]) + 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	var start, end, maxLength int
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			if s[i] == s[j] {
				if j - i == 1 || j - i == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	fmt.Println(dp)
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] {
				if maxLength < j-i {
					maxLength = j-i
					start = i
					end = j
				}
			}
		}
	}
	return s[start:end+1]
}

func longestPalindrome2(s string) string {
	var maxLength int
	var res string
	// 中心扩散算法
	for i := 0; i < len(s); i++ {
		l1, r1 := helper(i, i, s)
		l2, r2 := helper(i, i+1, s)
		len1, len2 := r1 - l1, r2 - l2
		if len1 > len2 {
			if maxLength < len1 {
				maxLength = len1
				res = s[l1:r1+1]
			}
		} else if maxLength < len2 {
			maxLength = len2
			res = s[l2:r2+1]
		}
	}
	return res
}

func helper(l, r int, s string) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l --
		r ++
	}
	return l+1, r-1
}