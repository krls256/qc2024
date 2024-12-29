package main

import (
	"fmt"
	"qc/bb84"
)

func main() {
	alice := bb84.Sender{}
	bob := bb84.Receiver{}

	alice.Init()
	bob.Init()

	fmt.Println("ALICE STATE\t", alice.StateString())
	fmt.Println("ALICE BASIS\t", alice.BasisString())
	fmt.Println("BOB BASIS\t", bob.BasisString())

	msgSent := alice.QuantumMessage()
	fmt.Println("ALICE MESSAGE (WE CAN'T SEE IT IN REAL LIFE)", msgSent.ToString())

	bob.AcceptMessage(msgSent)

	aSS := alice.SharedSecret(bob.Basis())
	bSS := bob.SharedSecret(alice.Basis())

	aSSBts := bb84.RoundSharedSecret(aSS)
	bSSBts := bb84.RoundSharedSecret(bSS)

	if len(aSS) <= 16 {
		fmt.Println("ALICE SHARED SECRET", aSS)
		fmt.Println("ALICE SHARED SECRET", bSS)

	} else {
		fmt.Println("ALICE SHARED SECRET", aSSBts)
		fmt.Println("ALICE SHARED SECRET", bSSBts)
	}

	fmt.Println("KeySize:", len(aSSBts))
}
