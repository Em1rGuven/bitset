# Bitset

A simple Bitset implementation in Go using `uint64` slices.

## Methods

| Method | Description |
|---|---|
| `NewBitset()` | Creates an empty Bitset |
| `SetBit(value)` | Sets a bit to 1 |
| `ClearBit(value)` | Sets a bit to 0 |
| `ToggleBit(value)` | Flips a bit |
| `IsBitOpen(value)` | Checks whether a bit is set |
| `CountSetBits()` | Returns the number of set bits |
| `Clear()` | Resets all bits to 0 |

- Dynamically grows when setting or toggling bits.
- The maximum supported bit index is `16,777,215`.
- Memory is allocated dynamically as needed.
- The implementation is not concurrency-safe.
