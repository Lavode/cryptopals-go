package cryptopals

import "fmt"

// Launcher allows registering and launching specific challenges.
type Launcher struct {
	challenges map[uint]map[uint]Challenge
}

// NewLauncher initializes a new launcher.
func NewLauncher() Launcher {
	launcher := Launcher{}
	launcher.challenges = make(map[uint]map[uint]Challenge)

	return launcher
}

// Register adds a new challenge to the launcher.
//
// Returns an error if a challenge for the given challenge set and ID already
// exist.
func (launcher *Launcher) Register(challenge Challenge) error {
	set := launcher.challengeSet(challenge.Set)

	_, exists := set[challenge.Challenge]
	if exists {
		return fmt.Errorf("Challenge set %d already contains challenge with ID %d", challenge.Set, challenge.Challenge)
	}

	set[challenge.Challenge] = challenge
	return nil
}

// challengeSet gets the challenge set with the given ID.
//
// If it does not exist yet, it is created.
func (launcher *Launcher) challengeSet(id uint) map[uint]Challenge {
	set, ok := launcher.challenges[id]
	if !ok {
		set = make(map[uint]Challenge)
		launcher.challenges[id] = set
	}

	return set
}

// Challenge returns the challenge with the given ID and set ID.
//
// The second return value indicates whether such a challenge exists.
func (launcher *Launcher) Challenge(setID, challengeID uint) (Challenge, bool) {
	set := launcher.challengeSet(setID)
	challenge, exists := set[challengeID]

	return challenge, exists
}
