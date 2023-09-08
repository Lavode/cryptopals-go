package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	// Equal inputs have distance 0
	a := []byte("YELLOW SUBMARINE")
	dist, err := HammingDistance(a, a)
	assert.NoError(t, err)
	assert.Equal(t, dist, uint(0))

	// Different inputs
	b := []byte("this is a test")
	c := []byte("wokka wokka!!!")
	dist, err = HammingDistance(b, c)
	assert.NoError(t, err)
	assert.Equal(t, dist, uint(37))
}

func TestHammingDistanceWithDifferentLengthInputs(t *testing.T) {
	a := []byte("YELLOW SUBMARINE")
	b := []byte("HELLO WORLD")

	_, err := HammingDistance(a, b)
	assert.Error(t, err)
}
