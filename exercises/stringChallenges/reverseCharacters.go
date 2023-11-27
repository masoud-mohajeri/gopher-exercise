package stringChallenges

import "fmt"

func ReverseCharacters(input string) string {
	inputLen := len(input)

	result := make([]rune, inputLen)

	for index, char := range input {
		targetIndex := inputLen - index - 1
		result[targetIndex] = char
	}
	fmt.Println(string(result))
	return string(result)

}
