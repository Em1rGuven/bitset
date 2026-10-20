package main

import (
	"math/bits"
)

const (
	BITS_PER_MASK int = 64
	MAX_VALUE     int = (1 << 24) - 1
)

type Bitset []uint64

func NewBitset() *Bitset {
	return &Bitset{}
}

func (b *Bitset) grow(value int) {
	index := value / BITS_PER_MASK
	if index < len(*b) {
		return
	}
	newLen := index + 1
	if newLen <= cap(*b) {
		*b = (*b)[:newLen]
		return
	}
	newCap := max(cap(*b)*2, newLen, 8)
	newSlice := make([]uint64, newLen, newCap)
	copy(newSlice, *b)
	*b = newSlice
}

func (b *Bitset) ToggleBit(value int) {
	if value < 0 || value > MAX_VALUE {
		return
	}
	index := value / BITS_PER_MASK
	bit := (value % BITS_PER_MASK)
	b.grow(value)
	(*b)[index] = (*b)[index] ^ (1 << bit)
}

func (b *Bitset) SetBit(value int) {
	if value < 0 || value > MAX_VALUE {
		return
	}
	index := value / BITS_PER_MASK
	bit := (value % BITS_PER_MASK)
	b.grow(value)
	(*b)[index] = (*b)[index] | (1 << bit)
}

func (b *Bitset) ClearBit(value int) {
	if value < 0 || value > MAX_VALUE {
		return
	}
	index := value / BITS_PER_MASK
	bit := (value % BITS_PER_MASK)
	if index >= len(*b) {
		return
	}
	(*b)[index] = (*b)[index] & ^(1 << bit)
}

func (b *Bitset) Clear() {
	clear(*b)
}

func (b *Bitset) IsBitOpen(value int) bool {
	if value < 0 || value > MAX_VALUE {
		return false
	}
	index := value / BITS_PER_MASK
	bit := (value % BITS_PER_MASK)
	if index >= len(*b) {
		return false
	}
	return ((*b)[index]>>bit)&1 == 1
}

func (b *Bitset) CountSetBits() int {
	r := 0
	for _, m := range *b {
		r += bits.OnesCount64(m)
	}
	return r
}
