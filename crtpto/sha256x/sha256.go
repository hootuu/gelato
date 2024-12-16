package sha256x

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

func SHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func SHA256Bytes(byteData []byte) string {
	hash := sha256.Sum256(byteData)
	return hex.EncodeToString(hash[:])
}

func IsSHA256(str string) bool {
	if len(str) != 64 {
		return false
	}

	matched, err := regexp.MatchString("^[0-9a-fA-F]{64}$", str)
	if err != nil {
		return false
	}
	return matched
}
