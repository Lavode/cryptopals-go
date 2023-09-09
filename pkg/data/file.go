package data

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// DataDir defines the name of the directory within which challenge data is stored.
const DataDir = "data"

// Base64FromFile reads base64-encoded data from a challenge file.
func Base64FromFile(set, challenge uint) ([]byte, error) {
	fh, err := os.Open(fmt.Sprintf("%s/%d_%d.txt", DataDir, set, challenge))
	if err != nil {
		return []byte{}, err
	}
	defer fh.Close()

	decoder := base64.NewDecoder(base64.StdEncoding, fh)
	out, err := io.ReadAll(decoder)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

// HexLines interprets the input file as newline-separated byte arrays, each of
// which is encoded as a hexadecimal string.
func HexLines(set, challenge uint) ([][]byte, error) {
	out := make([][]byte, 0)

	fh, err := os.Open(fmt.Sprintf("%s/%d_%d.txt", DataDir, set, challenge))
	if err != nil {
		return out, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		bytes, err := hex.DecodeString(scanner.Text())
		if err != nil {
			return out, err
		}

		out = append(out, bytes)
	}

	return out, nil
}
