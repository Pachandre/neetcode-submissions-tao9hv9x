func isPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	t := 0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			letters = append(letters, unicode.ToLower(r))
			t++
		}
	}
	n := len(letters)
	for i := range n / 2 {
		if letters[i] != letters[n-i-1] {
			return false
		}
	}
	return true
}