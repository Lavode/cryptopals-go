package logic

import "fmt"

// Xor calculates the logical XOR of two bytes slices of equal length.
//
// This function will panic if the slices are of different length.
func Xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Got slices of different length: %d != %d", len(a), len(b)))
	}
	out := make([]byte, len(a))

	for i := 0; i < len(a); i += 1 {
		out[i] = a[i] ^ b[i]
	}

	return out
}
