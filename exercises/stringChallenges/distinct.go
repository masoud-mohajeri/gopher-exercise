package stringChallenges

import "fmt"

func DistinctCharacterCount(input string) {

	charListMap := make(map[rune]bool)
	for _, char := range input {
		charListMap[char] = true
	}

	distinctChars := []rune{}
	for char := range charListMap {
		distinctChars = append(distinctChars, char)
	}

	fmt.Println(len(distinctChars))
}
