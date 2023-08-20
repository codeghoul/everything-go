package datastructures

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_BloomFilter(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		estimatedValueCount uint32
		filterSize          uint32
	}{
		{estimatedValueCount: 1000, filterSize: 8000},
		{estimatedValueCount: 5000, filterSize: 40000},
		{estimatedValueCount: 10000, filterSize: 80000},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			bf := NewBloomFilter(tc.filterSize, tc.estimatedValueCount)

			// Adding random strings to the Bloom filter
			for i := 0; i < int(tc.estimatedValueCount); i++ {
				bf.Add(uuid.NewString())
			}

			falsePositives := 0
			totalChecked := 10000

			// Checking for false positives
			for i := 0; i < totalChecked; i++ {
				if bf.Exists(uuid.NewString()) {
					falsePositives++
				}
			}

			actualFalsePositiveRate := float64(falsePositives) / float64(totalChecked)
			estimatedFalsePositiveRate := math.Pow(1-math.Exp(-float64(len(bf.hashFunctions))*float64(tc.estimatedValueCount)/float64(tc.filterSize)), float64(len(bf.hashFunctions)))

			assertionMargin := 0.004
			absDifference := math.Abs(estimatedFalsePositiveRate - actualFalsePositiveRate)

			// Assert that the absolute difference is within the allowed margin
			assert.LessOrEqualf(t, absDifference, assertionMargin,
				"Max difference between %v and %v allowed is %v, but difference was %v",
				estimatedFalsePositiveRate, actualFalsePositiveRate, assertionMargin, absDifference)
		})
	}
}
