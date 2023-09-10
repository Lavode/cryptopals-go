package analysis

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsECB(t *testing.T) {
	// Encryption of 'YELLOW SUBMARINE0123456789ABCDEFYELLOW SUBMARINE'
	ctxt, err := base64.StdEncoding.DecodeString("0apPZXiSZUL7tt2HbNIFCGzsN2A0R1nEk/OzZFF0CVrRqk9leJJlQvu23Yds0gUI")
	if err != nil {
		panic(err)
	}
	out, err := IsECB(ctxt)
	assert.NoError(t, err)
	assert.True(t, out)

	// Encryption of 'REDISH SUBMARINE0123456789ABCDEFYELLOW SUBMARINE'
	ctxt, err = base64.StdEncoding.DecodeString("XnUcaH6Cd1eq5Vh6Kp9QBmzsN2A0R1nEk/OzZFF0CVrRqk9leJJlQvu23Yds0gUI")
	if err != nil {
		panic(err)
	}
	out, err = IsECB(ctxt)
	assert.NoError(t, err)
	assert.False(t, out)
}
