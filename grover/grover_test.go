package grover

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// go test ./grover
func TestGrover(t *testing.T) {
	for size := 3; size <= 6; size++ {
		for secretNumber := uint(0); secretNumber < 1<<size; secretNumber++ {
			g := New(size, NewSecretFunc(secretNumber))

			require.Equal(t, secretNumber, g.Solve())
		}
	}
}
