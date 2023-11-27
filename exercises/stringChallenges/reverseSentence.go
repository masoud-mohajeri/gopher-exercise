package stringChallenges

import (
	"slices"
	"strings"
)

func ReverseSentence(input string) string {
	inputAsArray := SplitSentence(input)
	//numberOfWords := len(inputAsArray)
	//resultSlice := make([]string, numberOfWords)
	//for index, word := range inputAsArray {
	//	resultSlice[numberOfWords-index-1] = word
	//}
	//result := strings.Join(resultSlice, " ")

	// go 1.21
	slices.Reverse(inputAsArray)
	result := strings.Join(inputAsArray, " ")
	return result

}
