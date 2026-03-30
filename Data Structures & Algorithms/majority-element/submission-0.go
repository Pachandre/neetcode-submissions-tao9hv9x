func majorityElement(nums []int) int {
    c := map[int]int{}
	n := len(nums) / 2
	for _, v := range nums {
		c[v]++
		if c[v] > n {
			return v
		}
	}
	return -1
}
