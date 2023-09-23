package service

import (
	"fizz-buzz/library"
)

func Fizzbuzz(string1, string2 string, int1, int2, limit int) []string {
	var result string
	tabResult := make([]string, 0)

	for i := 1; i <= limit; i++ {

		result = ""
		if i%int1 == 0 {
			result += string1
		}
		if i%int2 == 0 {
			result += string2
		}
		if result == "" {
			result = library.IntToString(i)
		}

		tabResult = append(tabResult, result)
	}

	return tabResult
}
