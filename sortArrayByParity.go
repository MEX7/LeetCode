package leetcode

// 905. 按奇偶排序数组
func SortArrayByParity(A []int) []int {
	c := 0
	for i := range A {
		if A[i]%2 == 0 {
			A[c], A[i] = A[i], A[c]
			c++
		}
	}
	return A
}
