package leetcode

// 416. 分割等和子集
func CanPartition(nums []int) bool {
	if nums == nil {
		return true
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	sum = sum / 2

	n := len(nums)
	dp := make([][]bool, sum+1)

	for i := 0; i <= sum; i++ {
		dp[i] = make([]bool, n)
		if i == nums[0] {
			dp[i][0] = true
		}
	}

	for i := 0; i <= sum; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1]
			if i >= nums[j] {
				dp[i][j] = dp[i][j-1] || dp[i-nums[j]][j-1]
			}
		}
	}

	return dp[sum][n-1]
}
