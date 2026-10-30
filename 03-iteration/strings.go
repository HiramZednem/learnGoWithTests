package iteration

import "strings"

func IsWordInside(sentence, word string) string {

	isInside := strings.Contains(sentence, word)
	if (isInside) {
		return "yes"
	} 

	return "no"
}

func CountChars(sentence, char string) int {
	return strings.Count(sentence, char);
}

func CutPrefix(sentence, prefix string) string {
	wordCutted, found := strings.CutPrefix(sentence, prefix)

	if !found {
		wordCutted = "";
	}

	return wordCutted
}

func SeparateString(sentence, separator string) []string {
	if strings.TrimSpace(separator) == "" {
		return strings.Fields(sentence);
	}
	return strings.Split(sentence, separator)
}