package analysis

import (
	"github.com/Lavode/cryptopals-go/pkg/language"
	"github.com/Lavode/cryptopals-go/pkg/logic"
)

// XorFrequencyAnalysis attempts to find the single-byte key of a many-time-pad
// by using frequency analysis.
//
// It returns its guess at a message and the corresponding key as well as the
// resulting Hellinger distance.
func XorFrequencyAnalysis(ciphertext []byte) ([]byte, byte, float64) {
	// Hellinger distance is in [0, 1] so we can safely initialize this to 1
	minDistance := 1.0
	var bestKey byte
	var bestMessage []byte

	for keyCandidate := 0; keyCandidate < 256; keyCandidate++ {
		msgCandidate := logic.RepeatingXor(ciphertext, []byte{byte(keyCandidate)})
		frequencies := language.Normalize(language.FrequencyCount(string(msgCandidate)))
		dist := language.HellingerDistance(language.EnglishLetterFrequencies, frequencies)

		if dist < minDistance {
			minDistance = dist
			bestKey = byte(keyCandidate)
			bestMessage = msgCandidate
		}
	}

	return bestMessage, bestKey, minDistance
}
