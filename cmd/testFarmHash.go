package main

import (
	"fmt"
	"github.com/leemcloughlin/gofarmhash"
	"github.com/singularian/mdhash/farmHash32"
	"github.com/singularian/mdhash/farmHash64"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }
	fh32 := farmHash32.New(0, 4)
	fh32.Write(a)
	fmt.Printf("fh32 hex %x\n", fh32.Sum(nil))

	fh32.Write(bytes)
	fmt.Printf("fh32 hex %x\n", fh32.Sum(nil))

	hash := farmhash.Hash32(bytes)
	fmt.Printf("Hash32 (%s) is %x\n", str, hash)

	fh64 := farmHash64.New(0, 8, 234324324)
	fh64.Write(bytes)
	fmt.Printf("Hash64 (%s) is %x %d\n", str, fh64.Sum(nil), len(fh64.Sum(nil)))

	fh64.Write(a)
	fmt.Printf("Hash64 (%s) is %x %d\n", str, fh64.Sum(nil), len(fh64.Sum(nil)))

	hash64 := farmhash.Hash64WithSeed(bytes, 234324324)
	fmt.Printf("Hash64 (%s) is %x\n", str, hash64)
}

