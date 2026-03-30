func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
	for i, v := range nums {
		left := target - v
		if j, ok := seen[left]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return []int{}
}
