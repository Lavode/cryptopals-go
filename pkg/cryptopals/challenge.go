package cryptopals

import "log"

type Challenge struct {
	Set       uint
	Challenge uint
	Exec      func() error
}

func (c *Challenge) Run() {
	log.Printf("== Challenge %d.%d ==", c.Set, c.Challenge)
	c.Exec()
}
