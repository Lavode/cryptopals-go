package logic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitAlternating(t *testing.T) {
	input := make([]byte, 20)
	for i := byte(0); i < 20; i++ {
		input[i] = i
	}

	out := SplitAlternating(input, 3)

	assert.Equal(t, len(out), 3)
	assert.Equal(t, out[0], []byte{0, 3, 6, 9, 12, 15, 18})
	assert.Equal(t, out[1], []byte{1, 4, 7, 10, 13, 16, 19})
	assert.Equal(t, out[2], []byte{2, 5, 8, 11, 14, 17})
}

func ExampleSplitAlternating() {
	input := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	output := SplitAlternating(input, 3)

	fmt.Printf("[%v, %v, %v]", output[0], output[1], output[2])
	// Output: [[0 3 6], [1 4 7], [2 5]]
}

func TestMergeAlternating(t *testing.T) {
	input := [][]byte{
		{0, 4, 8, 12, 16},
		{1, 5, 9, 13},
		{2, 6, 10, 14, 17},
		{3, 7, 11, 15},
	}
	output := MergeAlternating(input)

	assert.Equal(t, len(output), 18)
	assert.Equal(t, output, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17})
}

func ExampleTestMergeAlternating() {
	input := [][]byte{
		{0, 1, 2, 3, 4},
		{5, 6},
		{7, 8, 9},
	}
	output := MergeAlternating(input)

	fmt.Printf("%v", output)
	// Output: [0 5 7 1 6 8 2 9 3 4]
}
