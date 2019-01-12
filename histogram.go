package main

import (
	"math/bits"
)

type memHucket struct {
	cnt uint64
	sum uint64
}

type memHistogram struct {
	cnt     uint64
	total   uint64
	max     uint64
	buckets [60]memHucket
}

func (h *memHistogram) add(n uint64) {
	index := bits.Len64(n) - 1
	if index < 0 {
		index = 0
	}
	bucket := &h.buckets[index]
	bucket.cnt++
	bucket.sum += n
	h.cnt++
	h.total += n

	if n > h.max {
		h.max = n
	}
}

func (h *memHistogram) accumSum(start uint64, end uint64) uint64 {
	var sum uint64
	for i := start; i <= end; i++ {
		sum += h.buckets[i].sum
	}
	return sum
}

func (h *memHistogram) accumCnt(start uint64, end uint64) uint64 {
	var sum uint64
	for i := start; i <= end; i++ {
		sum += h.buckets[i].cnt
	}
	return sum
}

type countHistogram struct {
	cnt     uint64
	total   uint64
	max     uint64
	buckets [41]uint64
}

func (h *countHistogram) add(n uint64) {
	index := bits.Len64(n) - 1
	h.buckets[index]++
	h.cnt++
	h.total += n

	if n > h.max {
		h.max = n
	}
}
