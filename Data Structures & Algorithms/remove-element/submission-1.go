func removeElement(nums []int, val int) int {
    t := 0
	for i := range nums {
		if nums[i] != val {
			nums[t] = nums[i]
			t++
		}
	}
	return t
}
