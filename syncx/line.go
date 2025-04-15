package syncx

import (
	"github.com/hootuu/gelato/errors"
	"sync"
)

type Line struct {
	mu sync.Mutex
}

func NewLine() *Line {
	return &Line{}
}

func (line *Line) Do(call func() *errors.Error) *errors.Error {
	line.mu.Lock()
	defer line.mu.Unlock()
	return call()
}
