package hexx

import (
	"encoding/hex"
	"github.com/hootuu/gelato/errors"
)

func Encode(src []byte) string {
	return hex.EncodeToString(src)
}

func Decode(hexStr string) ([]byte, *errors.Error) {
	dst, nE := hex.DecodeString(hexStr)
	if nE != nil {
		return nil, errors.System("hex.Decode failed", nE)
	}
	return dst, nil
}
