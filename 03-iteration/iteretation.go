package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repetead strings.Builder;
	for i := 0; i < repeatCount; i++ {
		repetead.WriteString(character)
	}

	return repetead.String();
}