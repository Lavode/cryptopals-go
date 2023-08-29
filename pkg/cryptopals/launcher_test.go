package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	launcher := NewLauncher()

	c := Challenge{Set: 1, Challenge: 2, Exec: noopChallenge}
	assert.NoError(t, launcher.Register(c))
	assert.Len(t, launcher.challenges[1], 1)

	c = Challenge{Set: 1, Challenge: 3, Exec: noopChallenge}
	assert.NoError(t, launcher.Register(c))
	assert.Len(t, launcher.challenges[1], 2)

	c = Challenge{Set: 2, Challenge: 2, Exec: noopChallenge}
	assert.NoError(t, launcher.Register(c))
	assert.Len(t, launcher.challenges[2], 1)

	c = Challenge{Set: 1, Challenge: 2, Exec: noopChallenge}
	assert.Error(t, launcher.Register(c))
	assert.Len(t, launcher.challenges[1], 2)
}

func TestChallenge(t *testing.T) {
	launcher := NewLauncher()
	c1 := Challenge{Set: 1, Challenge: 2, Exec: noopChallenge}
	launcher.Register(c1)
	c2 := Challenge{Set: 2, Challenge: 3, Exec: noopChallenge}
	launcher.Register(c2)

	// We cannot compare Challenge instances as they contain a function
	// member, which testify "cannot compare for equality".
	_, exists := launcher.Challenge(1, 2)
	assert.True(t, exists)

	_, exists = launcher.Challenge(2, 3)
	assert.True(t, exists)

	_, exists = launcher.Challenge(4, 3)
	assert.False(t, exists)
}

func noopChallenge() error {
	return nil
}
