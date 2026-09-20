func isAnagram(s string, t string) bool {
	sCount := make(map[rune]int)
	for _, char := range s {
		sCount[char] += 1
	}

	for _, char := range t {
		sCount[char] -= 1
	}

	for _, v := range sCount {
		if v != 0 {
			return false
		}
	}

	return true
}
