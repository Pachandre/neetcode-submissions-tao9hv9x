func sortColors(nums []int) {
    counts := [3]int{0, 0, 0}
	for _, v := range nums {
		counts[v]++
	}
	idx := 0
	for i := range nums {
		for counts[idx] == 0 {
			idx++
		}
		nums[i] = idx
		counts[idx]--
	}
}
