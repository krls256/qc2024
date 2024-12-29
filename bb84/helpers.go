package bb84

import (
	"bytes"
	"fmt"
)

func bitsString(bts []byte) string {
	buf := bytes.Buffer{}

	for _, b := range bts {
		buf.WriteString(fmt.Sprintf("%08b", b))
	}

	return buf.String()
}

func bytesToBoolSlice(bts []byte) []bool {
	result := make([]bool, len(bts)*8)

	for i := 0; i < len(bts); i++ {
		for j := range 8 {
			result[i*8+j] = (bts[i] & (1 << (7 - j))) != 0
		}
	}

	return result
}

func boolToBytesSlice(slice []bool) []byte {
	if len(slice)%8 != 0 {
		panic("wrong bool size")
	}

	result := make([]byte, len(slice)/8)

	for k := 0; k < len(slice); k++ {
		i, j := k/8, k%8

		if slice[k] {
			result[i] = result[i] | (1 << (7 - j))
		}
	}

	return result
}

func RoundSharedSecret(ss []bool) []byte {
	desiredLen := len(ss) - (len(ss) % 8)

	return boolToBytesSlice(ss[:desiredLen])
}

func SharedSecret(b1, b2, state []byte) []bool {
	b1B := bytesToBoolSlice(b1)
	b2B := bytesToBoolSlice(b2)
	stateBool := bytesToBoolSlice(state)
	resultState := []bool{}

	for i := 0; i < len(b1B); i++ {
		if b1B[i] == b2B[i] {
			resultState = append(resultState, stateBool[i])
		}
	}

	return resultState
}
