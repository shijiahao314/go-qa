package utils

import "strconv"

func StringToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func Uint64ToString(val uint64) (string, error) {
	return strconv.FormatUint(val, 10), nil
}
