package symmetric

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECBDecrypt(t *testing.T) {
	ctxt, err := base64.StdEncoding.DecodeString("WfWPPIw+c5aw5YWzuDItOGBlgz+SVZIMASp0HWRjvrPhIdJ9buksSw+3g9lB8z0dGfMc/9lF4RI5mtsSeC8yQWHQPFjiUYqMEmKwZB+zSBGFd0dQD56V0/sm5iRkVcdg")
	if err != nil {
		panic(err)
	}

	key := []byte("YELLOW SUBMARINE")

	expected := []byte("This is a test message which happened to be a multiple of 16 bytes. What are the chances? Small!")
	msg, err := ECBDecrypt(key, ctxt)
	assert.NoError(t, err)
	assert.Equal(t, expected, msg)
}

func TestECBDecryptWithIncorrectInputLength(t *testing.T) {
	// Ciphertext too short
	key := []byte("YELLOW SUBMARINE")
	ctxt := []byte("0123456789ABCDE")
	_, err := ECBDecrypt(key, ctxt)
	assert.Error(t, err)

	// Ciphertext too long
	ctxt = []byte("0123456789ABCDEF0")
	_, err = ECBDecrypt(key, ctxt)
	assert.Error(t, err)

	// Key too short
	key = []byte("YELLOW SUBMARIN")
	ctxt = []byte("0123456789ABCDEF")
	_, err = ECBDecrypt(key, ctxt)
	assert.Error(t, err)

	// Key too long
	key = []byte("YELLOW SUBMARINEYELLOW")
	ctxt = []byte("0123456789ABCDEF")
	_, err = ECBDecrypt(key, ctxt)
	assert.Error(t, err)
}
