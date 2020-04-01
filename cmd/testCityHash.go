package main

import (
	"fmt"
	"github.com/singularian/mdhash/cityHash64"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

/*	sp32 := spooky32.New(0, 4, 123213)
	sp32.Write(a)
	fmt.Printf("spooky32 bytes %x hex %x\n", a, sp32.Sum(nil))

	sp32.Write(bytes)
	fmt.Printf("spooky32 hex %x hash %x\n", bytes, sp32.Sum(nil))

	hash := spooky.Hash32(bytes, 123213)
	fmt.Printf("Hash32 (%s) is %x\n", str, hash)
*/
	ct64 := cityHash64.New(0, 8, 234324324)
	ct64.Write(bytes)
	fmt.Printf("City hash64 test 1 (%s) is %x %d\n", str, ct64.Sum(nil), len(ct64.Sum(nil)))

	ct64.Write(a)
	fmt.Printf("City Hash64 test 2 bytes (%x) is %x %d\n", a, ct64.Sum(nil), len(ct64.Sum(nil)))

	// hash64 := spooky.Hash64(bytes, 234324324)
	// fmt.Printf("Hash64 (%s) is %x\n", str, hash64)
}

