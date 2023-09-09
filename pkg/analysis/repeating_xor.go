package analysis

import (
	"fmt"
	"log"

	"github.com/Lavode/cryptopals-go/pkg/logic"
)

// MinRepeatingXorKeylength specifies the minimum key length of a
// many-times-pad cipher.
const MinRepeatingXorKeylength = 2

// MaxRepeatingXorKeylength specifies the maximum key length of a
// many-times-pad cipher.
const MaxRepeatingXorKeylength = 50

// RepeatingXorBlockCount specifies how many blocks to use to
// estimate the length of the key of a many-times-pad cipher.
const RepeatingXorBlockCount = 8

// RepeatingXor attempts to analyze a many-times-pad (repeating-xor) ciphertext.
//
// It returns the resulting plaintext and key.
func RepeatingXor(ciphertext []byte) ([]byte, []byte, error) {
	plaintext := make([]byte, 0)
	key := make([]byte, 0)

	keyLen, err := estimateRepeatingXorKeylength(ciphertext)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	log.Printf("Most likely keylength = %d", keyLen)

	// Split into smaller blocks of alternating bytes
	alternatingCtxt := logic.SplitAlternating(ciphertext, keyLen)

	// Analyze each block - all bytes of which are encrypted using the same key - in isolation.
	alternatingMsg := make([][]byte, keyLen)
	for i := 0; i < len(alternatingCtxt); i++ {
		messageChunk, keyByte, _ := XorFrequencyAnalysis(alternatingCtxt[i])
		alternatingMsg[i] = messageChunk

		key = append(key, keyByte)
	}

	// And merge back together in proper order
	plaintext = logic.MergeAlternating(alternatingMsg)

	return plaintext, key, nil
}

// estimateRepeatingXorKeylength attempts to find the key length used in a
// many-time-pad encryption.
//
// If the input is too short an error is returned.
func estimateRepeatingXorKeylength(ciphertext []byte) (int, error) {
	var bestKeyLen int
	// At most two blocks differ in every bit so we start with that distance
	var lowestDistance float64 = MaxRepeatingXorKeylength * 8

	for keyLen := MinRepeatingXorKeylength; keyLen <= MaxRepeatingXorKeylength; keyLen++ {
		requiredLength := keyLen * RepeatingXorBlockCount
		if len(ciphertext) < requiredLength {
			return 0, fmt.Errorf("Ciphertext too short. Need %d bytes but got only %d", requiredLength, len(ciphertext))
		}

		blocks := make([][]byte, RepeatingXorBlockCount)
		for i := 0; i < RepeatingXorBlockCount; i++ {
			blocks[i] = ciphertext[i*keyLen : (i+1)*keyLen]
		}

		var distance float64
		// We calculate the average hamming distance between any two of the blocks
		for i := 0; i < RepeatingXorBlockCount; i++ {
			for j := (i + 1); j < RepeatingXorBlockCount; j++ {
				d, err := logic.HammingDistance(blocks[i], blocks[j])
				if err != nil {
					return 0, err
				}

				// Normalize by block length
				distance += float64(d) / float64(keyLen)
			}
		}
		// With `n` blocks there are `n \times (n-1) / 2` pairs of two distinct blocks.
		averageDistance := distance / (RepeatingXorBlockCount * (RepeatingXorBlockCount - 1) / 2)

		if averageDistance < lowestDistance {
			lowestDistance = averageDistance
			bestKeyLen = keyLen
		}
	}

	return bestKeyLen, nil
}
