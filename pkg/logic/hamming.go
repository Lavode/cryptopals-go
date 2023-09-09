package logic

import (
	"fmt"
)

// HammingDistance calculates the hamming distance between two inputs of equal length.
//
// The hamming distance is the number of differing bits between the two inputs.
//
// If inputs differ in length, an error is returned.
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Inputs were of unequal length: %d != %d", len(a), len(b))
	}

	var distance int

	for i := 0; i < len(a); i++ {
		difference := a[i] ^ b[i]

		for difference > 0 {
			distance += int(difference & 0x01)
			difference = difference >> 1
		}
	}

	return distance, nil
}
