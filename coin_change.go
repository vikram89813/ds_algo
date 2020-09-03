package main

import "fmt"

var dp [][]int

func getDpArray(m int, n int) [][]int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	return dp
}

func count(arr []int, m int, n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if m <= 0 && n >= 1 {
		return 0
	}

	if dp[m][n] != 0 {
		return dp[m][n]
	}

	dp[m][n] = count(arr, m-1, n) + count(arr, m, n-arr[m-1])
	return dp[m][n]
}

func main() {
	n := 4
	arr := []int{1, 2, 3}
	dp = getDpArray(4, 5)
	fmt.Println(count(arr, len(arr), n))
}
