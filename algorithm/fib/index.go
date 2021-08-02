package fib


func fib(n int) int {
	var dp = []int {0, 1}
	for i := 2; i <= n;i++ {
		dp = append(dp,dp[0] + dp[1])
		dp = dp[1:]
	}
	return dp[1]
}