package bloom_filter

import (
	"bufio"
	"fmt"
	"os"
)

type HashFunc func(string) uint

type BloomFilter struct {
	byteMap   []byte
	bitCount  uint
	hashFuncs []HashFunc
}

func NewBloomFilter(bitLen int, fn ...HashFunc) *BloomFilter {
	if len(fn) == 0 {
		panic("NewBloomFilter: at least one hash function is required")
	}

	size := bitLen / 8
	bf := &BloomFilter{
		byteMap:   make([]byte, size),
		bitCount:  uint(bitLen),
		hashFuncs: fn,
	}

	return bf
}

func (bf *BloomFilter) Add(element string) {
	// bs := []byte(element)
	for _, fn := range bf.hashFuncs {
		index := fn(element) % bf.bitCount
		bf.setBit(index)
	}
}

func (bf *BloomFilter) Exists(element string) bool {
	for _, fn := range bf.hashFuncs {
		index := fn(element) % bf.bitCount
		if !bf.getBit(index) {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) Dump(file string) error {
	fout, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		return fmt.Errorf("Dump: failed to create file %s: %w", file, err)
	}

	defer fout.Close()

	w := bufio.NewWriter(fout)
	if _, err := w.WriteString(fmt.Sprintf("bitCount: %d\n", bf.bitCount)); err != nil {
		return fmt.Errorf("Dump: failed to write metadata: %w", err)
	}

	for i := 0; i < len(bf.byteMap); i += 4 {
		end := i + 4
		if end > len(bf.byteMap) {
			end = len(bf.byteMap)
		}

		line := bf.byteMap[i:end]
		if _, err := w.WriteString(fmt.Sprintf("%08d: %08b\n", i, line)); err != nil {
			return fmt.Errorf("Dump: failed to write byteMap: %w", err)
		}
	}

	return w.Flush()
}

func (bf *BloomFilter) getBit(index uint) bool {
	var i uint = index / 8
	var b uint = uint(bf.byteMap[i])
	var target uint = 1 << (index % 8)

	return (b & target) == target
}

func (bf *BloomFilter) setBit(index uint) {
	var i uint = index / 8
	var target uint = 1 << (index % 8)
	bf.byteMap[i] |= byte(target)
}
