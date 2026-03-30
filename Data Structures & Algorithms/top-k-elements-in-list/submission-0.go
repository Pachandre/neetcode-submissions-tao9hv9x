func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	maxEl := func() int {
		maxK := -1
		maxV := -1
		for k, v := range mp {
			if v > maxV {
				maxK = k
				maxV = v
			}
		}
		return maxK
	}
	res := make([]int, k)
	for i := range k {
		maxK := maxEl()
		delete(mp, maxK)
		res[i] = maxK
	}
	return res
}
