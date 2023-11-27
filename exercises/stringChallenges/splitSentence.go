package stringChallenges

func SplitSentence(input string) []string {
	word := []rune{}
	result := []string{}
	for index, char := range input {
		if char != ' ' {
			word = append(word, char)
		}

		if char == ' ' || index == len(input)-1 {
			result = append(result, string(word))
			word = []rune{}
			continue
		}
	}
	return result
}
