package bloom_filter

import (
	"crypto/sha256"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	hash1 := DefaultHashFunc(1204)
	hash2 := DefaultHashFunc(666)
	hash3 := DefaultHashFunc(19)

	bf := NewBloomFilter(100, hash1, hash2, hash3)
	a, b, c, d := "昔人已乘黃鶴去", "此地空餘黃鶴樓", "黃鶴一去不復返", "白雲千載空悠悠"

	bf.Add(a)
	bf.Add(b)

	assert.Equal(t, bf.Exists(c), false)
	assert.Equal(t, bf.Exists(d), false)
	assert.Equal(t, bf.Exists(a), true)
	assert.Equal(t, bf.Exists(b), true)

	assert.Nil(t, bf.Dump("./bloom_dump.bin"))

	bf2 := NewBloomFilter(70, hash1, hash2, hash3)
	assert.Error(t, bf2.Load("./bloom_dump.bin"))

	bf2 = NewBloomFilter(100, hash1, hash2, hash3)
	assert.Nil(t, bf2.Load("./bloom_dump.bin"))

	assert.Equal(t, bf2.Exists(c), false)
	assert.Equal(t, bf2.Exists(d), false)
	assert.Equal(t, bf2.Exists(a), true)
	assert.Equal(t, bf2.Exists(b), true)

	assert.Nil(t, bf2.Dump("./bloom_dump2.bin"))
	chk1, err := calculateFileHash("./bloom_dump.bin")
	assert.Nil(t, err)
	chk2, err := calculateFileHash("./bloom_dump2.bin")
	assert.Nil(t, err)
	assert.Equal(t, string(chk1), string(chk2))
}

func calculateFileHash(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
