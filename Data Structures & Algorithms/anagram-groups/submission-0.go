type Key = [26]int
func groupAnagrams(strs []string) [][]string {
	mp := map[Key][]string{}
	for _, str := range strs {
		key := Key{}
		for _, r := range str {
			key[int(r - 'a')]++
		}
		mp[key] = append(mp[key], str)
	}
	res := make([][]string, 0, len(mp))
	for _, strs := range mp {
		res = append(res, strs)
	}
	return res
}
