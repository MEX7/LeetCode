package leetcode

// 55. 跳跃游戏
func CanJump(nums []int) bool {
	l := len(nums)
	if l < 1 {
		return false
	}
	if l == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}

	zero := -1
	for i := l - 2; i >= 0; i-- {
		if zero > 0 {
			if i+nums[i] > zero {
				zero = -1
			}
			continue
		}

		if nums[i] == 0 {
			zero = i
			continue
		}
	}
	if zero < 0 {
		return true
	}
	return false
}
