package piscine

func CountA(words []string) int {
	count := 0
	for _, word := range words {
		if len(word) > 0 && word[0] == 'a' {
			count = count + 1
		}
	}
	return count
}

func main() {
}
