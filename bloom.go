package bloom_filter

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
