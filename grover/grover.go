package grover

import (
	"github.com/itsubaki/q"
	"github.com/samber/lo"

	"math"
	"math/bits"
)

func New(size int, secretFunc func([]q.Qubit) []q.Qubit) Grover {
	return Grover{size: size, secretFunc: secretFunc}
}

type Grover struct {
	size       int
	secretFunc func([]q.Qubit) []q.Qubit
}

func (g Grover) Solve() uint {
	simulator := q.New()
	allQubits := []q.Qubit{}

	for range g.size {
		allQubits = append(allQubits, simulator.Zero())
	}

	// Part 1 on my doc image
	simulator.H(allQubits...)
	iterations := int(math.Pi / 4 * math.Sqrt(math.Pow(2, float64(bits.Len(uint(g.size))))))

	for i := 0; i < iterations; i++ {
		step(allQubits, simulator, g.secretFunc)
	}

	measure := simulator.State()
	maxProb := 0.0
	maxValue := uint(0)

	for _, state := range measure {
		if state.Probability > maxProb {
			maxProb = state.Probability
			maxValue = uint(state.Int[0])
		}
	}

	return maxValue
}

func step(allQubits []q.Qubit, simulator *q.Q, secretFunc func([]q.Qubit) []q.Qubit) {
	qLast := allQubits[len(allQubits)-1]

	// Part 2 on my doc image
	secQubits := secretFunc(allQubits)
	if len(secQubits) > 0 {
		simulator.X(secQubits...)
	}

	simulator.H(qLast).
		ControlledNot(allQubits[:len(allQubits)-1], qLast).
		H(qLast)

	if len(secQubits) > 0 {
		simulator.X(secQubits...)
	}

	// Part 3 on my doc image
	simulator.
		H(allQubits...).
		X(allQubits...).
		H(qLast).
		ControlledNot(allQubits[:len(allQubits)-1], qLast).
		H(qLast).
		X(allQubits...).
		H(allQubits...)
}

func NewSecretFunc(secret uint) func([]q.Qubit) []q.Qubit {
	return func(allQubits []q.Qubit) []q.Qubit {
		mask := uint(1)
		toDelete := []int{}

		for i := 0; i < bits.Len(secret); i++ {
			if mask&secret != 0 {
				toDelete = append(toDelete, len(allQubits)-1-i)
			}

			mask = mask << 1
		}

		return lo.Filter(allQubits, func(item q.Qubit, i int) bool {
			return !lo.Contains(toDelete, i)
		})
	}
}
