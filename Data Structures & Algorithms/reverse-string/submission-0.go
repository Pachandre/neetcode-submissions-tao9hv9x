func reverseString(s []byte) {
	n := len(s)
	for i := range n / 2 {
		j := n - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
