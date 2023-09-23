package library

import "strconv"

func IsInt(value string) (bool, int) {
	intValue, err := strconv.Atoi(value)
	bResult := (err == nil)

	return bResult, intValue
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}
