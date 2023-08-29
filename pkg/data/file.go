package data

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const DATA_DIR = "data"

// HexLines interprets the input file as newline-separated byte arrays, each of
// which is encoded as a hexadecimal string.
func HexLines(set, challenge uint) ([][]byte, error) {
	out := make([][]byte, 0)

	fh, err := os.Open(fmt.Sprintf("%s/%d_%d.txt", DATA_DIR, set, challenge))
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
