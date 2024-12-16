package rand

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/hootuu/gelato/errors"
	"math/big"
	"strings"
)

func Int64() (int64, *errors.Error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 0, errors.System("crypto.rand failed to generate random number")
	}
	return int64(binary.BigEndian.Uint64(b[:])), nil
}

func String(length int) (string, *errors.Error) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+`-=,./?;:'"
	var result strings.Builder
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", errors.System("crypto.rand failed to generate random number")
		}
		result.WriteByte(charset[index.Int64()])
	}

	return result.String(), nil
}
