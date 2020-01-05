package main

import (
	"fmt"
	// "encoding/hex"
	"github.com/singularian/mdhash/poly1305"
)

func main() {

	a := []byte{ 0, 10, 22, 38, 240, 171, 146, 123 }

	key := [32]byte{0xda, 0x84, 0xbc, 0xab, 0x02, 0x67, 0x6c, 0x38, 0xcd, 0xb0, 0x15, 0x60, 0x42, 0x74, 0xc2, 0xaa}
        var keyx [32]byte
        copy(keyx[:], key[:32])

	p := poly1305.New(key)
	p.Write(a)


	fmt.Printf("poly1305 bytes hex %x length %d\n", p.Sum(nil), len(p.Sum(nil)))

	aaa := []byte{ 0, 10, 22, 38, 240, 171, 146, 123, 1, 1, 1, 1 }
	p.Write(aaa)
	fmt.Printf("poly1305 bytes hex %x length %d\n", p.Sum(nil), len(p.Sum(nil)))
}

