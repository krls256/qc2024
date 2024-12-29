package bb84

import (
	"bytes"
	"fmt"
	"github.com/itsubaki/q"
)

func NewMessage(size int) Message {
	msg := make(Message, 0, size)

	for range size {
		qsim := q.New()
		qsim.Zero()

		msg = append(msg, qsim)
	}

	return msg
}

type Message []*q.Q

func (msg Message) ToString() string {
	buf := bytes.Buffer{}

	for i := 0; i < len(msg); i++ {
		amp := msg[i].Amplitude()

		if inDelta(real(amp[0]), 1, 0.01) {
			buf.WriteString("|0>")

			continue
		}

		if inDelta(real(amp[1]), 1, 0.01) {
			buf.WriteString("|1>")

			continue
		}

		if inDelta(real(amp[0]), 0.707106781, 0.01) && inDelta(real(amp[1]), 0.707106781, 0.01) {
			buf.WriteString("|+>")

			continue
		}

		if inDelta(real(amp[0]), 0.707106781, 0.01) && inDelta(real(amp[1]), -0.707106781, 0.01) {
			buf.WriteString("|->")

			continue
		}

		panic(fmt.Sprintf("unexpected amp value %v (qubit: %v)", amp, i+1))
	}

	return buf.String()
}

func inDelta(value, cmp, delta float64) bool {
	left, right := cmp-delta, cmp+delta
	return value > left && value < right
}
