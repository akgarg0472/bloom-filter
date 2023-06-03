package main

import (
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

var hashFunction hash.Hash32

type BloomFilter struct {
	filter   []uint8
	capacity int
}

// function to create a new bloom filter with a given capacity
func NewBloomFilter(capacity int) *BloomFilter {
	if capacity <= 0 {
		panic("Bloom filterapacity must be greater than 0")
	}

	hashFunction = murmur3.New32WithSeed(uint32(time.Now().UnixMilli()))

	return &BloomFilter{
		filter:   make([]uint8, capacity),
		capacity: capacity,
	}
}

// function to check if the key exists in the filter and returns the index along with the result
func (b *BloomFilter) Exists(key string) (string, int, bool) {
	index := hashKey(key) % b.capacity
	arrayIndex := index / 8
	bitIndex := index % 8
	return key, arrayIndex, b.filter[arrayIndex]&(1<<bitIndex) == 1
}

// function to add a key to the filter
func (b *BloomFilter) Add(key string) {
	index := hashKey(key) % b.capacity
	arrayIndex := index / 8
	bitIndex := index % 8
	b.filter[arrayIndex] = b.filter[arrayIndex] | (1 << uint(bitIndex))
}

// function to hash the key to a number
func hashKey(key string) int {
	hashFunction.Write([]byte(key))
	var result = hashFunction.Sum32()
	hashFunction.Reset()
	return int(result)
}
