package maps

// 문자열 수
func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func hasDupeRune(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}

		runeSet[r] = true
	}

	return false
}

func hasDupeRune2(s string) bool {
	runeSet := map[rune]struct{}{}

	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}

		runeSet[r] = struct{}{}
	}

	return false
}
