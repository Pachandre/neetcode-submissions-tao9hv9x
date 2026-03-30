func twoSum(numbers []int, target int) []int {
	seen := map[int]int{}
	for i, v := range numbers {
		dif := target - v
		if j, ok := seen[dif]; ok {
			return []int{j + 1, i + 1}
		}
		seen[v] = i
	}
	return []int{0, 0}
}
