package lib

import "strings"

func ToSnakeCase(str string) string {
	runes := []rune(str)
	length := len(runes)

	var result []rune
	for i := 0; i < length; i++ {
		if i > 0 && runes[i] >= 'A' && runes[i] <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, runes[i])
	}

	return strings.ToLower(string(result))
}
