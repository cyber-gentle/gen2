package gen2

import (
	"strings"
)

func UpperCase(s string) string {

	char := []rune(s)
	for i := 0; i < len(char); i++ {
		if char[i] >= 90 && char[i] <= 122 {
			char[i] -= 32
		}
	}
	return string(char)
}

func LowerCase(s string) string {

	var newChar = make([]rune, 0)

	allChar := []rune(s)
	for _, c := range allChar {
		if c >= 65 && c <= 90 {
			c += 32
			newChar = append(newChar, c)
		} else {
			newChar = append(newChar, c)
		}
	}
	return string(newChar)
}

func Capitalize(s string) string {
	var newWord = make([]string, 0)

	words := strings.Fields(s)
	for _, word := range words {
		char := []rune(word)
		if char[0] >= 97 && char[0] <= 122 {
			char[0] -= 32
		}
		word = string(char)
		newWord = append(newWord, word)
	}
	return strings.Join(newWord, " ")
}

func CapitaliseWord(s string) string {
	var newWord = make([]string, 0)

	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		char := []rune(words[i])
		if char[0] >= 90 && char[0] <= 122 {
			char[0] -= 32
		}
		newWord = append(newWord, string(char))
	}

	return strings.Join(newWord, " ")
}
