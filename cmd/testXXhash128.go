package main

import (
	"fmt"
	// "encoding/hex"
	"github.com/singularian/mdhash/xxhash_128"
)

func main() {

	a := []byte{ 0, 10, 22, 38, 240, 171, 146, 123 }

	xx := xxhash_128.New()
	xx.Write(a)
	fmt.Printf("xx bytes hex %x length %d\n", xx.Sum(nil), len(xx.Sum(nil)))
	xx.Reset()

}

