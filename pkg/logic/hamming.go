package logic

import "fmt"

func HammingDistance(a, b []byte) (uint, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Inputs were of unequal length: %d != %d", len(a), len(b))
	}

	var distance uint = 0

	for i := 0; i < len(a); i += 1 {
		difference := a[i] ^ b[i]

		for difference > 0 {
			distance += uint(difference & 0x01)
			difference = difference >> 1
		}
	}

	return distance, nil
}
