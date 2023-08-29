package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {
	out := Xor([]byte{0x00, 0x27, 0x45}, []byte{0xAF, 0x45, 0x20})

	assert.Equal(t, out, []byte{0xAF, 0x62, 0x65})
}

func TestXorPanics(t *testing.T) {
	assert.Panics(t, func() { Xor([]byte{0x00, 0x27, 0x45}, []byte{0xAF, 0x45}) })
	assert.Panics(t, func() { Xor([]byte{0x00, 0x27, 0x45}, []byte{0xAF, 0x45, 0x20, 0x21}) })
}