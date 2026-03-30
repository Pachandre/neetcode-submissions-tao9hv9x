func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n + 1)
	left[0] = 1
	right := make([]int, n + 1)
	right[n] = 1
	for i := range n {
		left[i + 1] = left[i] * nums[i]
		right[n - i - 1] = right[n - i] * nums[n - i - 1]
	}
	res := make([]int, n)
	for i := range n {
		res[i] = left[i] * right[i + 1]
	}
	return res
}
