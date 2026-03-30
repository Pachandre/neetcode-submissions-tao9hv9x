func longestConsecutive(nums []int) int {
	s := map[int]bool{}
	for _, v := range nums {
		s[v] = true
	}
	maxLen := 0
	seen := map[int]int{}
	for num := range s {
		if s[num - 1] { continue }
		c := num
		for s[c + 1] {
			c++
			if v, ok := seen[c]; ok {
				c = v
				break
			}
		}
		seen[num] = c
		maxLen = max(maxLen, c - num + 1)
	}
	return maxLen
}