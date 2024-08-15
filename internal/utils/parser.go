package utils

import "strconv"

func ParseUint(s string) uint {
	// Parse the string into a uint
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}

	return uint(i)
}
