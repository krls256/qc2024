package bb84

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bytesToBoolSlice(t *testing.T) {
	bts := []byte{0b11110000, 0b00110011}

	booleans := []bool{
		true, true, true, true, false, false, false, false,
		false, false, true, true, false, false, true, true,
	}

	require.Equal(t, booleans, bytesToBoolSlice(bts))
}

func Test_boolToBytesSlice(t *testing.T) {
	bts := []byte{0b11110000, 0b00110011}

	booleans := []bool{
		true, true, true, true, false, false, false, false,
		false, false, true, true, false, false, true, true,
	}

	require.Equal(t, bts, boolToBytesSlice(booleans))
}
