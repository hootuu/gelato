package syncx

import (
	"github.com/hootuu/gelato/errors"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Single struct {
	mu      sync.RWMutex
	muCache *cache.Cache
}

func NewSingle() *Single {
	return &Single{
		mu:      sync.RWMutex{},
		muCache: cache.New(50*time.Second, 200*time.Second),
	}
}

func (s *Single) Do(key string, call func() *errors.Error) *errors.Error {
	s.mu.RLock()
	var singleMu *sync.Mutex
	muObj, found := s.muCache.Get(key)
	if !found {
		s.mu.RUnlock()
		s.mu.Lock()
		singleMu = &sync.Mutex{}
		s.muCache.SetDefault(key, singleMu)
		s.mu.Unlock()
	} else {
		singleMu = muObj.(*sync.Mutex)
		s.mu.RUnlock()
	}
	singleMu.Lock()
	defer singleMu.Unlock()
	return call()
}
