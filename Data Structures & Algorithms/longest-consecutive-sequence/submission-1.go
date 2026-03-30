func longestConsecutive(nums []int) int {
	s := map[int]bool{}
	for _, v := range nums {
		s[v] = true
	}
	maxLen := 0
	for num := range s {
		c := num
		for s[c + 1] {
			c++
		}
		maxLen = max(maxLen, c - num + 1)
	}
	return maxLen
}