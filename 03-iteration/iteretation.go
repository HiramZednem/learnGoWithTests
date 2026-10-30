package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repetead strings.Builder;
	for i := 0; i < repeatCount; i++ {
		repetead.WriteString(character)
	}

	return repetead.String();
}