package utils

import "strconv"

func ATOIwithDeafult(a string, defaultVal int) int {
	val, err := strconv.Atoi(a)
	if err != nil && val < 1 {
		return defaultVal
	}

	return val
}
