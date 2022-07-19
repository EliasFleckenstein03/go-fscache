package fscache

import (
	"github.com/gofrs/flock"
)

// Locks the file that backs the given key.
//
// If the call is successful, it's the caller's responsibility to call Unlock on the returned lock.
func (cd *CacheDir) Lock(key ...CacheKey) (*flock.Flock, error) {
	l := flock.New(cd.cachePath(key...))
	return l, l.Lock()
}
