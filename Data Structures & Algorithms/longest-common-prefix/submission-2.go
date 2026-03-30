func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := 200
	data := make([][]byte, len(strs))
	for i, str := range strs {
		data[i] = []byte(str)
		minLen = min(minLen, len(data[i]))
	}

	for idx := 0; idx < minLen; idx++ {
		for _, row := range data {
			if row[idx] != data[0][idx] {
				return string(data[0][:idx])
			}
		}
	}
	return string(data[0][:minLen])
}
