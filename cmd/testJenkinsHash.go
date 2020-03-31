package main

import (
	"fmt"
	"github.com/tildeleb/hashland/jenkins"
	"github.com/singularian/mdhash/jenkins32"
	"github.com/singularian/mdhash/jenkins64"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

	jh32 := jenkins32.New(0, 4, 123213)
	jh32.Write(a)
	fmt.Printf("fh32 bytes %x hex %x\n", a, jh32.Sum(nil))

 	jh32.Write(bytes)
	fmt.Printf("jh32 hex %x hash %x\n", bytes, jh32.Sum(nil))

	hash := jenkins.Hash232(bytes, 123213)
	fmt.Printf("Hash32 (%s) is %x\n", str, hash)

	jh64 := jenkins64.New(0, 8, 234324324)
	jh64.Write(bytes)
	fmt.Printf("Jenkins hash64 test 1 (%s) is %x %d\n", str, jh64.Sum(nil), len(jh64.Sum(nil)))

	jh64.Write(a)
	fmt.Printf("Jenkins Hash64 test 2 bytes (%x) is %x %d\n", a, jh64.Sum(nil), len(jh64.Sum(nil)))

	hash64 := jenkins.Hash264(bytes, 234324324)
	fmt.Printf("Hash64 (%s) is %x\n", str, hash64)
}

