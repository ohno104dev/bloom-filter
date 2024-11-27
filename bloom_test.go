package bloom_filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	hash1 := DefaultHashFunc(1204)
	hash2 := DefaultHashFunc(666)
	hash3 := DefaultHashFunc(19)

	bf := NewBloomFilter(1<<20, hash1, hash2, hash3)
	a, b, c, d := "昔人已乘黃鶴去", "此地空餘黃鶴樓", "黃鶴一去不復返", "白雲千載空悠悠"

	bf.Add(a)
	bf.Add(b)

	assert.Equal(t, bf.Exists(c), false)
	assert.Equal(t, bf.Exists(d), false)
	assert.Equal(t, bf.Exists(a), true)
	assert.Equal(t, bf.Exists(b), true)
}
