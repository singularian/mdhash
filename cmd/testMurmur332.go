package main

import (
	"fmt"
	"github.com/singularian/mdhash/murmur332"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

	m332 := murmur332.New(0, 4, 123213)
	m332.Write(a)
	fmt.Printf("murmur3 32 bytes %x hex %x\n", a, m332.Sum(nil))

	m332.Write(bytes)
	fmt.Printf("murmur 3 32 hex %x hash %x\n", bytes, m332.Sum(nil))

/*	hash := spooky.Hash32(bytes, 123213)
	fmt.Printf("Hash32 (%s) is %x\n", str, hash)

	sp64 := spooky64.New(0, 8, 234324324)
	sp64.Write(bytes)
	fmt.Printf("Spooky hash64 test 1 (%s) is %x %d\n", str, sp64.Sum(nil), len(sp64.Sum(nil)))

	sp64.Write(a)
	fmt.Printf("Spooky Hash64 test 2 bytes (%x) is %x %d\n", a, sp64.Sum(nil), len(sp64.Sum(nil)))

	hash64 := spooky.Hash64(bytes, 234324324)
	fmt.Printf("Hash64 (%s) is %x\n", str, hash64)
*/
}

