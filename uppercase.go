package piscine

func CountUppercase(words []string) int {
	count := 0
	for _, word := range words {
		if len(word) > 0 && word[0] >= 'A' && word[0] <= 'Z' {
			count = count + 1
		}
	}
	return count
}
