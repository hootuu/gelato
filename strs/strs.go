package strs

import (
	"strconv"
	"strings"
)

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func ToUint64(s string) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func ToInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return value
}
