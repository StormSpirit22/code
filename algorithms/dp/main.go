package main

import "fmt"

func main() {
	fmt.Println(minDistance("sea", "eat"))
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
