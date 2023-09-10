// main is the main package.
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Lavode/cryptopals-go/pkg/analysis"
	"github.com/Lavode/cryptopals-go/pkg/cryptopals"
	"github.com/Lavode/cryptopals-go/pkg/data"
	"github.com/Lavode/cryptopals-go/pkg/logic"
	"github.com/Lavode/cryptopals-go/pkg/symmetric"
)

func main() {
	launcher := cryptopals.NewLauncher()

	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 1, Exec: hexToBase64})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 2, Exec: fixedXor})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 3, Exec: singleByteXorCipher})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 4, Exec: detectSingleByteXor})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 5, Exec: repeatingByteXor})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 6, Exec: analyzeRepeatingByteXor})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 7, Exec: aesEcb})
	launcher.Register(cryptopals.Challenge{Set: 1, Challenge: 8, Exec: detectAesEcb})

	if len(os.Args) != 3 {
		usage()
	}

	setID, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("Invalid set ID. Must be numerical.")
		usage()
	}

	challengeID, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		usage()
		fmt.Println("Invalid challenge ID. Must be numerical.")
	}

	challenge, ok := launcher.Challenge(uint(setID), uint(challengeID))
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

// 1-3
func singleByteXorCipher() error {
	ctxt, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		panic(err)
	}

	msg, key, distance := analysis.XorFrequencyAnalysis(ctxt)
	log.Printf("Deduced plaintext = '%s' with key = %x (%c) (distance = %v)", msg, key, key, distance)

	return nil
}

// 1-4
func detectSingleByteXor() error {
	ctxts, err := data.HexLines(1, 4)
	if err != nil {
		panic(err)
	}

	var bestMessage []byte
	var bestKey byte
	lowestDistance := 1.0

	for _, ctxt := range ctxts {
		msg, key, dist := analysis.XorFrequencyAnalysis(ctxt)

		if dist < lowestDistance {
			bestMessage = msg
			bestKey = key
			lowestDistance = dist
		}
	}

	log.Printf("Found most-likely ciphertext and decrypted it to '%s' with key = %x (%c) (distance = %v)", bestMessage, bestKey, bestKey, lowestDistance)

	return nil
}

// 1-5
func repeatingByteXor() error {
	msg := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	expected, err := hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	if err != nil {
		panic(err)
	}

	ctxt := logic.RepeatingXor(msg, key)

	if !bytes.Equal(ctxt, expected) {
		return fmt.Errorf("Got unexpected result: %x", ctxt)
	}

	log.Printf("Got correct result")

	return nil
}

// 1-6
func analyzeRepeatingByteXor() error {
	ctxt, err := data.Base64FromFile(1, 6)
	if err != nil {
		panic(err)
	}

	msg, key, err := analysis.RepeatingXor(ctxt)
	if err != nil {
		return err
	}

	log.Printf("Message = %s", msg)
	log.Printf("Key = %s", key)

	return nil
}

// 1-7
func aesEcb() error {
	key := []byte("YELLOW SUBMARINE")
	ctxt, err := data.Base64FromFile(1, 7)
	if err != nil {
		panic(err)
	}

	msg, err := symmetric.ECBDecrypt(key, ctxt)
	if err != nil {
		return err
	}

	log.Printf("Plaintext: %s", msg)

	return nil
}

// 1-8
func detectAesEcb() error {
	ciphertexts, err := data.HexLines(1, 8)
	if err != nil {
		panic(err)
	}

	for _, ctxt := range ciphertexts {
		isECB, err := analysis.IsECB(ctxt)
		if err != nil {
			return err
		}

		if isECB {
			log.Printf("Found potential ECB ciphertext: %s", hex.EncodeToString(ctxt))
		}
	}

	return nil
}
