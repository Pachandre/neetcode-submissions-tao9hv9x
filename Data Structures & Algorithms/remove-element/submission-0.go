func removeElement(nums []int, val int) int {
    t := 0
	for i := range nums {
		nums[t] = nums[i]
		if nums[i] != val {
			t++
		}
	}
	return t
}
