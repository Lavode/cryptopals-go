package padding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkcs7Pad(t *testing.T) {
	// 11 bytes
	input := []byte("Hello world")

	padded, err := Pkcs7Pad(input, 12)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello world\x01"), padded)

	padded, err = Pkcs7Pad(input, 15)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello world\x04\x04\x04\x04"), padded)

	// Empty message
	padded, err = Pkcs7Pad([]byte{}, 4)
	assert.NoError(t, err)
	assert.Equal(t, []byte("\x04\x04\x04\x04"), padded)

	padded, err = Pkcs7Pad(input, 11+255)
	assert.NoError(t, err)
	assert.Equal(t, 11+255, len(padded))
	for i := 11; i < 11+255; i++ {
		assert.Equal(t, byte(0xFF), padded[i])
	}
}

func TestPkcs7PadInvalidLength(t *testing.T) {
	// 11 bytes
	input := []byte("Hello world")

	// Output length less than input length
	_, err := Pkcs7Pad(input, 8)
	assert.Error(t, err)

	// Output length equal to input length
	_, err = Pkcs7Pad(input, 11)
	assert.Error(t, err)

	// Output length exceeds input length + 255
	_, err = Pkcs7Pad(input, 11+256)
	assert.Error(t, err)
}

func TestPkcs7Unpad(t *testing.T) {
	// 4 bytes of padding
	padded := []byte("Hello world\x04\x04\x04\x04")
	msg, err := Pkcs7Unpad(padded)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello world"), msg)

	// Sanity check that we don't cut off too much
	padded = []byte("Hello world\x03\x03\x03\x03")
	msg, err = Pkcs7Unpad(padded)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello world\x03"), msg)

	// Empty message
	padded = []byte("\x03\x03\x03")
	msg, err = Pkcs7Unpad(padded)
	assert.NoError(t, err)
	assert.Equal(t, []byte{}, msg)

	// 255 bytes of padding
	padded = make([]byte, 255+11)
	copy(padded, []byte("Hello world"))
	for i := 11; i < 11+255; i++ {
		padded[i] = byte(0xFF)
	}
	msg, err = Pkcs7Unpad(padded)
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello world"), msg)
}

func TestPkcs7UnpadInvalidPadding(t *testing.T) {
	// Missing one byte of padding
	padded := []byte("Hello world\x04\x04\x04")
	_, err := Pkcs7Unpad(padded)
	assert.Error(t, err)

	// Missing multiple bytes of padding
	padded = []byte("Hello world\xFF")
	_, err = Pkcs7Unpad(padded)
	assert.Error(t, err)
}
