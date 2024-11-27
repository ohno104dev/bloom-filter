package bloom_filter

import (
	"github.com/dgryski/go-farm"
)

func DefaultHashFunc(seed uint32) HashFunc {
	return func(s string) uint {
		bs := []byte(s)
		return uint(farm.Hash32WithSeed(bs, seed))
	}
}
