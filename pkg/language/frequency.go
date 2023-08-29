package language

import "math"

// LetterHistogram is a map between individual letters and the frequency with
// which they occur in a text.
type LetterHistogram = map[rune]float64

// EnglishLetterFrequencies specifies the normalized frequencies of lowercase
// alphabetic letters and the whitespace character in English texts.
var EnglishLetterFrequencies = LetterHistogram{
	'a': 0.0653,
	'b': 0.012,
	'c': 0.0223,
	'd': 0.0343,
	'e': 0.1012,
	'f': 0.0175,
	'g': 0.0159,
	'h': 0.0486,
	'i': 0.0558,
	'j': 0.0012,
	'k': 0.0061,
	'l': 0.0319,
	'm': 0.0191,
	'n': 0.0534,
	'o': 0.0598,
	'p': 0.0151,
	'q': 0.0008,
	'r': 0.0478,
	's': 0.0502,
	't': 0.0725,
	'u': 0.0223,
	'v': 0.0078,
	'w': 0.0191,
	'x': 0.0012,
	'y': 0.0159,
	'z': 0.0006,
	' ': 0.2024,
}

// FrequencyCount counts the frequency of individual letters (runes) in the
// input string.
func FrequencyCount(input string) map[rune]uint {
	out := make(map[rune]uint)

	for _, c := range input {
		out[c] += 1
	}

	return out
}

// Normalize normalizes letter counts.
func Normalize(counts map[rune]uint) LetterHistogram {
	sum := uint(0)
	for _, count := range counts {
		sum += count
	}

	out := make(LetterHistogram)
	for char, count := range counts {
		out[char] = float64(count) / float64(sum)
	}

	return out
}

// HellingerDistance calculates the Hellinger distance between two histograms,
// a metric indicating how similar the two histograms are to each other.
//
// The distance is 0 IFF the two histograms are equal, and 1 IFF for every
// possibility where one histogram assigns a non-zero probability, the other
// assigns a zero probability.
func HellingerDistance(a, b LetterHistogram) float64 {
	dist := 0.0

	// We misuse this as a set to store the union of keys of a and b
	keys := make(map[rune]bool)
	for char := range a {
		keys[char] = true
	}
	for char := range b {
		keys[char] = true
	}

	// 1) SUM [ (sqrt(p_i) - sqrt(q_i))^2 ]
	for char := range keys {
		dist += math.Pow(math.Sqrt(a[char])-math.Sqrt(b[char]), 2)
	}

	// 2) 1/sqrt(2) * sqrt(...)
	dist = math.Sqrt(dist) / math.Sqrt(2)

	return dist
}
