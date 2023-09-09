package cryptopals

import "log"

// Challenge represents a single cryptopals.com challenge.
type Challenge struct {
	Set       uint
	Challenge uint
	Exec      func() error
}

// Run executes the challenge's code.
func (c *Challenge) Run() {
	log.Printf("== Challenge %d.%d ==", c.Set, c.Challenge)

	if err := c.Exec(); err != nil {
		log.Printf("Error during challenge execution: %v", err)
	}
}
