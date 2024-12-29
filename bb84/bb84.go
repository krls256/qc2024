package bb84

import (
	"crypto/rand"
)

// We hardcode state size like in protocol
//const KeySize = 1
const KeySize = 80

// Sender == Alice
type Sender struct {
	state []byte
	basis []byte
}

func (s *Sender) Init() {
	s.state = make([]byte, KeySize)
	s.basis = make([]byte, KeySize)

	if _, err := rand.Read(s.state); err != nil {
		panic(err)
	}

	if _, err := rand.Read(s.basis); err != nil {
		panic(err)
	}
}

func (s *Sender) StateString() string {
	return bitsString(s.state)
}

func (s *Sender) BasisString() string {
	return bitsString(s.basis)
}

func (s *Sender) QuantumMessage() Message {
	basisBool := bytesToBoolSlice(s.basis)
	stateBool := bytesToBoolSlice(s.state)

	msg := NewMessage(len(basisBool))

	for i := 0; i < len(msg); i++ {
		if stateBool[i] {
			msg[i].X()
		}
		if basisBool[i] {
			msg[i].H()
		}
	}

	return msg
}

func (s *Sender) Basis() []byte {
	return s.basis
}

func (s *Sender) SharedSecret(rBasis []byte) []bool {
	return SharedSecret(s.basis, rBasis, s.state)
}

// Receiver == Bob
type Receiver struct {
	basis []byte
	state []byte
}

func (r *Receiver) Init() {
	r.basis = make([]byte, KeySize)

	if _, err := rand.Read(r.basis); err != nil {
		panic(err)
	}
}

func (r *Receiver) BasisString() string {
	return bitsString(r.basis)
}

func (r *Receiver) AcceptMessage(msg Message) {
	basisBool := bytesToBoolSlice(r.basis)

	boolState := make([]bool, len(msg))

	for i := 0; i < len(msg); i++ {
		if basisBool[i] {
			msg[i].H()
		}

		if inDelta(real(msg[i].Measure().Amplitude()[1]), 1, 0.01) {
			boolState[i] = true
		}
	}

	r.state = boolToBytesSlice(boolState)
}

func (r *Receiver) Basis() []byte {
	return r.basis
}

func (r *Receiver) SharedSecret(sBasis []byte) []bool {
	return SharedSecret(r.basis, sBasis, r.state)
}
