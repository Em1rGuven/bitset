package main

import "testing"

func TestBitset(t *testing.T) {
	bitset := NewBitset()

	bitset.SetBit(43)
	if !bitset.IsBitOpen(43) {
		t.Error("bit did not open")
	}

	bitset.ClearBit(43)
	if bitset.IsBitOpen(43) {
		t.Error("bit did not close")
	}

	bitset.SetBit(65)
	bitset.SetBit(402)
	bitset.SetBit(99)
	if bitset.CountSetBits() != 3 {
		t.Error("set bits count is not true")
	}

	bitset.Clear()
	if bitset.CountSetBits() != 0 {
		t.Error("bitset did not reset")
	}

	bitset.SetBit(63)
	bitset.ToggleBit(63)
	if bitset.IsBitOpen(63) {
		t.Error("bit toggle error")
	}

	bitset.Clear()
	bitset.SetBit(1 << 29)
	bitset.SetBit(-1)
	if bitset.CountSetBits() > 0 {
		t.Error("invalid bit opened")
	}
}
