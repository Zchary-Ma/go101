package leetcode

import (
	"strings"
)

func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}
