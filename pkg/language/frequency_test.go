package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencyCount(t *testing.T) {
	msg := "hello world, this is your captain speaking"
	freqs := FrequencyCount(msg)
	expected := map[rune]uint{
		'h': 2,
		'e': 2,
		'l': 3,
		'o': 3,
		'w': 1,
		'r': 2,
		' ': 6,
		'd': 1,
		',': 1,
		't': 2,
		'i': 4,
		's': 3,
		'y': 1,
		'u': 1,
		'c': 1,
		'a': 3,
		'p': 2,
		'n': 2,
		'k': 1,
		'g': 1,
	}

	assert.Equal(t, freqs, expected)
}

func TestFrequencyCountCapitalization(t *testing.T) {
	msg := "HeLlo"
	freqs := FrequencyCount(msg)
	expected := map[rune]uint{
		'h': 1,
		'e': 1,
		'l': 2,
		'o': 1,
	}

	assert.Equal(t, freqs, expected)
}

func TestLetterHistogram(t *testing.T) {
	msg := "hello"
	freqs := Normalize(FrequencyCount(msg))
	expected := LetterHistogram{
		'h': 0.2,
		'e': 0.2,
		'l': 0.4,
		'o': 0.2,
	}

	for c, f := range expected {
		assert.InDelta(t, freqs[c], f, 1e-6)
	}
}

func TestHellingerDistance(t *testing.T) {
	a := LetterHistogram{
		'a': 0.3,
		'b': 0.2,
		'c': 0.5,
	}

	b := LetterHistogram{
		'a': 0.1,
		'c': 0.3,
		'd': 0.6,
	}

	dist := HellingerDistance(a, b)
	expected := 0.662945

	assert.InDelta(t, dist, expected, 1e-6)
}
