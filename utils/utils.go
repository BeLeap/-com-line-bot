package utils

import (
	"math/rand"
	"strings"
	"time"
)

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Select(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[Random(0, len(slice))]
}

func TrimWhitespace(content string) string {
	return strings.Trim(content, " ")
}

func TrimPrefix(content string, prefix string) string {
	result := TrimWhitespace(content)
	result = strings.TrimPrefix(result, prefix)
	return TrimWhitespace(content)
}

func SplitRealContent(content string, prefix string, delimeter string) []string {
	trimmedContent := TrimPrefix(content, prefix)
	return strings.Split(trimmedContent, delimeter)
}
