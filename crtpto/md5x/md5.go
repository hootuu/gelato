package md5x

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

func MD5(str string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func IsMD5(str string) bool {
	matched, err := regexp.MatchString("^[0-9a-fA-F]{32}$", str)
	if err != nil {
		return false
	}
	return matched
}
