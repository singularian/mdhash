package main

import (
	"fmt"
	"io"
	"encoding/hex"
	"github.com/singularian/mdhash/hw32"
)

func main() {
	key, err := hex.DecodeString("119102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") 
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}


	hw32 := hw32.New(0, 4, key)
	a := []byte{ 0, 10, 22, 38, 240, 171, 146, 123 }
	hw32.Write(a)
	fmt.Printf("hw32 hex %x\n", hw32.Sum(nil))

	teststring := "test hw32"
	hw32.Reset()
	io.WriteString(hw32, teststring)
	fmt.Printf("hw32 hex %x\n", hw32.Sum(nil))

}

