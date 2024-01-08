package utils

import (
	"strconv"
)

func StringToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func StringToUint(str string) (uint, error) {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func Uint64ToString(val uint64) (string, error) {
	return strconv.FormatUint(val, 10), nil
}
