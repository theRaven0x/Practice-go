package piscine

func FilterA(words []string) []string {
	result := []string{}
	for _, word := range words {
		if len(word) > 0 && word[0] == 'a' {
			result = append(result, word)
		}
	}
	return result
}
