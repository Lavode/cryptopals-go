package symmetric

import (
	"crypto/aes"
	"fmt"
)

// ECBDecrypt decrypts a ciphertext with AES in ECB mode.
func ECBDecrypt(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return []byte{}, fmt.Errorf("Key must be 16B, got %d", len(key))
	}

	if len(ciphertext)%aes.BlockSize!= 0 {
		return []byte{}, fmt.Errorf("Ciphertext must be multiple of %dB, got %d", aes.BlockSize, len(ciphertext))
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	plaintext := make([]byte, len(ciphertext))

	for block := 0; block < len(ciphertext)/aes.BlockSize; block++ {
		from := block * aes.BlockSize
		to := (block + 1) * aes.BlockSize

		cipher.Decrypt(plaintext[from:to], ciphertext[from:to])
	}

	return plaintext, nil
}
