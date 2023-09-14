package padding

import (
	"fmt"
)

// Pkcs7Pad pads the input message to the desired output length using PKCS#7 padding.
//
// At most 255 bytes of padding can be added, that is `length` must not exceed `len(msg) + 255`.
//
// If the output length is less or equal to the lenght of the message, or exceeds what can be done, an error is returned.
func Pkcs7Pad(msg []byte, length int) ([]byte, error) {
	padBytes := length - len(msg)

	if padBytes < 1 {
		return []byte{}, fmt.Errorf("Output length must be larger than input length")
	} else if padBytes > 255 {
		return []byte{}, fmt.Errorf("Output length may not exceed input length by more than 255 bytes")
	}

	padded := make([]byte, length)
	copy(padded, msg)
	for i := len(msg); i < length; i++ {
		padded[i] = byte(padBytes)
	}

	return padded, nil
}

// Pkcs7Unpad removes a PKCS#7 padding from the message.
//
// If the padding was invalid, an error is returned.
func Pkcs7Unpad(padded []byte) ([]byte, error) {
	padBytes := padded[len(padded)-1]
	messageLength := len(padded) - int(padBytes)

	if messageLength < 0 {
		return []byte{}, fmt.Errorf("Invalid padding")
	}

	for i := messageLength; i < len(padded); i++ {
		if padded[i] != padBytes {
			return []byte{}, fmt.Errorf("Invalid padding")
		}
	}

	return padded[0:messageLength], nil
}
