package util

import "strconv"

func StringToIntWithDefault(str string, d int) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		return d
	}

	return n
}
