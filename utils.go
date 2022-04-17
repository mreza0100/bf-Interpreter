package brainfuck

import (
	"regexp"
)

var regex = regexp.MustCompile(`[^\[\]<>+\-.,]`)

func trim(instructions byte) string {
	return regex.ReplaceAllString(string(instructions), "")
}
