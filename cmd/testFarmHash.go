package main

import (
	"fmt"
	"github.com/leemcloughlin/gofarmhash"
	"github.com/singularian/mdhash/farmHash32"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }
	fh32 := farmHash32.New(0, 4, a)
	fh32.Write(a)
	fmt.Printf("fh32 hex %x\n", fh32.Sum(nil))

	fh32.Write(bytes)
	fmt.Printf("fh32 hex %x\n", fh32.Sum(nil))

	hash := farmhash.Hash32(bytes)
	fmt.Printf("Hash32(%s) is %x\n", str, hash)

}

