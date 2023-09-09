package logic

import (
	"fmt"
)

// Xor calculates the logical XOR of two bytes slices of equal length.
//
// This function will panic if the slices are of different length.
func Xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Got slices of different length: %d != %d", len(a), len(b)))
	}

	return RepeatingXor(a, b)
}

// RepeatingXor calculates the logical XOR of two byte slices of arbitrary
// length. Access to the shorter of the two will wrap around automatically if
// required.
func RepeatingXor(a, b []byte) []byte {
	out := make([]byte, len(a))

	n := len(a)
	if len(b) > len(a) {
		n = len(b)
	}

	for i := 0; i < n; i++ {
		out[i] = a[i%len(a)] ^ b[i%len(b)]
	}

	return out
}
