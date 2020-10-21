package search

import "strings"

// Run поиск подстроки по значениям ассоциативного массива
func Run(substring string, data map[string]string) (result map[string]string) {
	result = make(map[string]string)

	for k, v := range data {
		if strings.Contains(v, substring) {
			result[k] = v
		}
	}
	return result
}
