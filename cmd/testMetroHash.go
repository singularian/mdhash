package main

import (
	"fmt"
	"github.com/singularian/mdhash/metro64"
)

func main() {

	str := "hello world"
	bytes := []byte(str)
	a := []byte{ 11, 10, 22, 38, 240, 171, 146, 123 }

	mt64 := metro64.New(0, 8, 1232123)
	mt64.Write(bytes)
	fmt.Printf("Metro hash64 test 1 (%s) is %x %d\n", str, mt64.Sum(nil), len(mt64.Sum(nil)))

	mt64.Write(a)
	fmt.Printf("Metro Hash64 test 2 bytes (%x) is %x %d\n", a, mt64.Sum(nil), len(mt64.Sum(nil)))

}

