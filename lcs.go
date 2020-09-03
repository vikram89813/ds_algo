package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func lcs(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	fmt.Println(lcs(X, Y))
}
