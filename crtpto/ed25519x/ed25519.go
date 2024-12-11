package ed25519x

import (
	"crypto/ed25519"
	"crypto/rand"
	"github.com/hootuu/gelato/errors"
)

func NewRandom() ([]byte, []byte, *errors.Error) {
	pub, pri, nErr := ed25519.GenerateKey(rand.Reader)
	if nErr != nil {
		return nil, nil, errors.System("gen ed25519 key fail", nErr)
	}
	return pub, pri, nil
}