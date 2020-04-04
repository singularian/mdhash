package main

import (
	"fmt"
	"github.com/singularian/mdhash/marvin32"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

	m332 := marvin32.New(0, 4, 123213)
	m332.Write(a)
	fmt.Printf("marvin 32 bytes %x hex %x\n", a, m332.Sum(nil))

	m332.Write(bytes)
	fmt.Printf("marvin 32 hex %x hash %x\n", bytes, m332.Sum(nil))

}

