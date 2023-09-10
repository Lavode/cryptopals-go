package analysis

import (
	"crypto/aes"
	"fmt"
)

// IsECB attempts to check whether a ciphertext might be the result of AES-ECB encryption.
//
// It does so by checking whether there are any duplicate blocks of ciphertext.
//
// An error is returned if the ciphertext is not a multiple of the AES block size.
func IsECB(ciphertext []byte) (bool, error) {
	if len(ciphertext)%aes.BlockSize != 0 {
		return false, fmt.Errorf("Ciphertext must be multiple of %dB, was %d", aes.BlockSize, len(ciphertext))
	}

	blockCount := make(map[[aes.BlockSize]byte]int)
	for blockIdx := 0; blockIdx < len(ciphertext)/aes.BlockSize; blockIdx++ {
		from := blockIdx * aes.BlockSize
		to := (blockIdx + 1) * aes.BlockSize

		var block [aes.BlockSize]byte
		copy(block[:], ciphertext[from:to])

		blockCount[block]++
	}

	for _, cnt := range blockCount {
		if cnt > 1 {
			return true, nil
		}
	}

	return false, nil
}
