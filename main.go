package main

import (
	"fmt"
	"qc/grover"
)

func main() {
	fmt.Println(grover.New(5, grover.NewSecretFunc(23)).Solve())
}
