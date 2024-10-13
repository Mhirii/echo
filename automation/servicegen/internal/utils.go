package internal

import "strings"

func isSnakeCase(str string) bool {
	// only _ and a-z
	for i, r := range str {
		if i == 0 && r == '_' {
			return false
		}
		if r == '_' {
			continue
		}
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func snakeToPascal(str string) string {
	underscore := false
	var res string
	for i, r := range str {
		if i == 0 {
			res += strings.ToUpper(string(r))
			continue
		}
		if underscore {
			res += string(r - 32)
			underscore = false
			continue
		}
		if r == '_' {
			underscore = true
			continue
		}
		res += string(r)
	}
	return res
}

func capitalize(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
