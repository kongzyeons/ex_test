package utils

import (
	"regexp"
	"strings"
)

func GroupTextWord(input string) map[string]int {
	results := map[string]int{}

	pattern := "[.,]"
	regexpObj, _ := regexp.Compile(pattern)
	// if err != nil {
	// 	fmt.Println("Error compiling regex:", err)
	// 	return
	// }
	wordLow := strings.ToLower(input)
	wordRegex := regexpObj.ReplaceAllString(wordLow, "")
	words := strings.Fields(wordRegex)
	for i := range words {
		if words[i] != "" {
			results[words[i]]++
		}
	}
	return results
}
