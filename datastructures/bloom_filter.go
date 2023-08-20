package datastructures

import (
	"hash"
	"math"
	"math/rand"
	"time"

	"github.com/spaolacci/murmur3"
)

// getHash calculates the index for the filter using the provided hash function, key, and filter size.
func getHash(hasher hash.Hash32, key string, size uint32) uint32 {
	hasher.Write([]byte(key))
	result := hasher.Sum32() % size
	hasher.Reset()
	return result
}

// BloomFilter represents a probabilistic data structure that efficiently tests set membership.
type BloomFilter struct {
	size          uint32        // Size of the filter (number of bits)
	filter        *BitArray     // Actual filter data (you should define or import BitArray type)
	hashFunctions []hash.Hash32 // Hash functions for generating hash values
}

// NewBloomFilter creates a new Bloom filter instance with the specified size and estimated value count.
func NewBloomFilter(size uint32, estimatedValueCount uint32) *BloomFilter {
	// Calculate the optimal number of hash functions
	hashFuncCount := optimalHashFunctions(size, estimatedValueCount)
	rand.Seed(time.Now().UnixNano())

	// Create hash functions with random seeds
	hashFunctions := make([]hash.Hash32, hashFuncCount)
	for i := 0; i < hashFuncCount; i++ {
		hashFunctions[i] = murmur3.New32WithSeed(rand.Uint32())
	}

	return &BloomFilter{
		size:          size,
		filter:        NewBitArray(size), // You should create or import NewBitArray function
		hashFunctions: hashFunctions,
	}
}

// Add adds a new key to the Bloom filter by setting the corresponding bits in the filter.
func (bf *BloomFilter) Add(key string) {
	for _, fn := range bf.hashFunctions {
		idx := getHash(fn, key, bf.size)
		bf.filter.Set(idx)
	}
}

// Exists checks if a key might exist in the Bloom filter. It returns true if the key is potentially present,
// and false if it is definitely not present.
func (bf *BloomFilter) Exists(key string) bool {
	for _, fn := range bf.hashFunctions {
		idx := getHash(fn, key, bf.size)
		if !bf.filter.Get(idx) {
			return false
		}
	}

	return true
}

// optimalHashFunctions calculates the optimal number of hash functions using the formula: k = (m / n) * ln(2)
func optimalHashFunctions(size uint32, estimatedValueCount uint32) int {
	optimalK := int(math.Ceil(float64(size) / float64(estimatedValueCount) * math.Log(2)))
	return optimalK
}
