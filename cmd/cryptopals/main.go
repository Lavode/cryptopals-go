package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Lavode/cryptopals-go/pkg/cryptopals"
	"github.com/Lavode/cryptopals-go/pkg/logic"
)

func main() {
	launcher := cryptopals.NewLauncher()

	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 1, Exec: hexToBase64})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 2, Exec: fixedXor})

	if len(os.Args) != 3 {
		usage()
	}

	setId, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("Invalid set ID. Must be numerical.")
		usage()
	}

	challengeId, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		usage()
		fmt.Println("Invalid challenge ID. Must be numerical.")
	}

	challenge, ok := launcher.Challenge(uint(setId), uint(challengeId))
	if !ok {
		fmt.Println("No such challenge found")
		usage()
	}

	challenge.Run()
}

func usage() {
	fmt.Println("Usage: ./cryptopals <set_id> <challenge_id>")
	os.Exit(2)
}

// 1-1
func hexToBase64() error {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	bytes := make([]byte, len(input)/2)
	_, err := hex.Decode(bytes, input)
	if err != nil {
		return fmt.Errorf("Unable to hex-decode input: %v", err)
	}

	output := base64.StdEncoding.EncodeToString(bytes)
	log.Printf("Resulting base64 encoding: %s\n", output)

	return nil
}

// 1-2
func fixedXor() error {
	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		panic(err)
	}

	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		panic(err)
	}

	c := logic.Xor(a, b)
	log.Printf("%x XOR %x = %x", a, b, c)
	if hex.EncodeToString(c) != "746865206b696420646f6e277420706c6179" {
		panic("Incorrect output")
	}

	return nil
}
