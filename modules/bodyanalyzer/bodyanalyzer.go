package bodyanalyzer

import (
	"encoding/hex"
	"hash"
)

// Analyzer IO Writer making it easy to get some stats from a reader
type Analyzer struct {
	hasher        hash.Hash
	shouldGetSize bool
	size          int64
}

// New create a body analyzer with specified config
func New(hasher hash.Hash, getSize bool) *Analyzer {
	return &Analyzer{
		hasher:        hasher,
		shouldGetSize: getSize,
	}
}

func (b *Analyzer) Write(p []byte) (n int, err error) {
	csize := len(p)
	b.size += int64(csize)

	if b.hasher != nil {
		b.hasher.Write(p)
	}

	return csize, nil
}

// Size Get current size read in the Analyzer
func (b *Analyzer) Size() int64 {
	return b.size
}

// HexHash Get current hash
func (b *Analyzer) HexHash() string {
	return hex.EncodeToString(b.hasher.Sum(nil))
}
