package utils

import "strconv"

func IntsToStrings(arr []int) []string {
	res := make([]string, len(arr))
	for i, n := range arr {
		res[i] = strconv.Itoa(n)
	}

	return res
}
