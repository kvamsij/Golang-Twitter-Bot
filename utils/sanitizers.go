package utils

import (
	"regexp"
	"strings"
)

func RemoveUrlsAndHashTagsFromString(str string) string {
	index := 0
	words := strings.Split(str, " ")

	for _, word := range words {

		if !isUrl(word) && !hasHashTagorIsUser(word) {
			words[index] = word
			index++
		}
	}

	words = words[:index]
	textWithAlphaNumericChar := removeNonAlphanumericCharacters(strings.Join(words, " "))
	sanitizedWord := removeSingleCharacter(textWithAlphaNumericChar)
	return sanitizedWord
}

func removeSingleCharacter(text string) string {
	words := strings.Split(text, " ")
	index := 0
	for _, word := range words {

		if len(word) > 1 {
			words[index] = word
			index++
		}
	}
	words = words[:index]
	return strings.Join(words, " ")
}

func isUrl(word string) bool {

	if strings.HasPrefix(word, "https://") || strings.HasPrefix(word, "http://") {
		return true
	}
	return false
}

func hasHashTagorIsUser(word string) bool {

	if strings.HasPrefix(word, "#") || strings.HasPrefix(word, "@") {
		return true
	}
	return false
}

func removeNonAlphanumericCharacters(text string) string {
	reg, err := regexp.Compile("[^a-zA-z]+")
	if err != nil {
		panic(err)
	}
	return reg.ReplaceAllString(text, " ")
}
