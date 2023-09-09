package logic

// SplitAlternating splits a byte slice into `n` slices, with input bytes alternating between output slices.
func SplitAlternating(in []byte, n int) [][]byte {
	out := make([][]byte, n)
	for i := 0; i < n; i++ {
		out[i] = make([]byte, 0)
	}

	for idx, b := range in {
		chunkIdx := idx % n

		out[chunkIdx] = append(out[chunkIdx], b)
	}

	return out
}

// MergeAlternating merges bytes slices, with output bytes being pulled from the input slices in alternating order.
func MergeAlternating(in [][]byte) []byte {
	count := 0
	maxLength := 0
	for _, slice := range in {
		count += len(slice)
		if maxLength < len(slice) {
			maxLength = len(slice)
		}
	}

	out := make([]byte, 0)

	chunkIdx := 0
	offsets := make([]int, len(in))
	for len(out) < count {
		if len(in[chunkIdx]) > offsets[chunkIdx] {
			// Can pull another byte from this slice
			out = append(out, in[chunkIdx][offsets[chunkIdx]])
			offsets[chunkIdx]++
		}

		chunkIdx = (chunkIdx + 1) % len(in)
	}

	return out
}
