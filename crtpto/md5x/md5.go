package md5x

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
)

func New(prefix string, split string, paras ...string) string {
	if len(paras) == 0 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(prefix))
	for _, s := range paras {
		h.Write([]byte(split))
		h.Write([]byte(s))
	}
	return hex.EncodeToString(h.Sum(nil))
}

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
