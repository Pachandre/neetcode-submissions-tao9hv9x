func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp1 := map[rune]int{}
	mp2 := map[rune]int{}
	for _, v := range s {
		mp1[v]++
	}
	for _, v := range t {
		mp2[v]++
	}
	if len(mp1) != len(mp2) {
		return false
	}
	for r, c := range mp1 {
		if mp2[r] != c {
			return false
		}
	}
	return true
}
