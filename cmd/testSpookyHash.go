package main

import (
	"fmt"
	"github.com/tildeleb/hashland/spooky"
	"github.com/singularian/mdhash/spooky32"
	"github.com/singularian/mdhash/spooky64"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

	sp32 := spooky32.New(0, 4, 123213)
	sp32.Write(a)
	fmt.Printf("spooky32 bytes %x hex %x\n", a, sp32.Sum(nil))

	sp32.Write(bytes)
	fmt.Printf("spooky32 hex %x hash %x\n", bytes, sp32.Sum(nil))

	hash := spooky.Hash32(bytes, 123213)
	fmt.Printf("Hash32 (%s) is %x\n", str, hash)

	sp64 := spooky64.New(0, 8, 234324324)
	sp64.Write(bytes)
	fmt.Printf("Spooky hash64 test 1 (%s) is %x %d\n", str, sp64.Sum(nil), len(sp64.Sum(nil)))

	sp64.Write(a)
	fmt.Printf("Spooky Hash64 test 2 bytes (%x) is %x %d\n", a, sp64.Sum(nil), len(sp64.Sum(nil)))

	hash64 := spooky.Hash64(bytes, 234324324)
	fmt.Printf("Hash64 (%s) is %x\n", str, hash64)
}

