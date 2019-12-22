package main

import (
	"fmt"
	"hash/fnv"
	// "encoding/hex"
	"github.com/skeeto/cubehash"
	"github.com/OneOfOne/xxhash"
	_ "github.com/martinlindhe/gogost"
	"github.com/martinlindhe/gogost/gost34112012256"
//	"github.com/aead/poly1305"
	"golang.org/x/crypto/poly1305"
//	"github.com/skeeto/chacha-go"
)


func main() {

	// teststring := "test of the SHA1 128 hash"
	a := []byte{ 0, 10, 22, 38, 240, 171, 146, 123 }
	c := cubehash.New()
	c.Write(a)
	fmt.Printf("sha1282 bytes hex % x\n", c.Sum(nil))
	fmt.Printf("cubehash %x %d\n", c.Sum(nil), len(c.Sum(nil)))

	fn := fnv.New128()
	fn.Write(a)
	fmt.Printf("fnv128 hex % x\n", fn.Sum(nil))
	

	fnx := fnv.New128a()
	fnx.Write(a)
	fmt.Printf("fnv128a hex %x %d\n", fnx.Sum(nil),  len(fnx.Sum(nil)))


	xx := xxhash.New64()
	// xx := xxhash.New()
	xx.Write(a)
	fmt.Printf("xxhash hex %x %d\n", xx.Sum(nil),  len(xx.Sum(nil)))

	gost := gost34112012256.New()
	gost.Write(a)
	fmt.Printf("gost hex %x %d\n", gost.Sum(nil),  len(gost.Sum(nil)))

	// key, _ := hex.DecodeString("000102030405060708090a0b0cff0e0f101112131415161718191a1b1c1d1e1f")
	key := [32]byte{0xda, 0x84, 0xbc, 0xab, 0x02, 0x67, 0x6c, 0x38, 0xcd, 0xb0, 0x15, 0x60, 0x42, 0x74, 0xc2, 0xaa}
	var keyx [32]byte
	copy(keyx[:], key[:32])
	// poly := poly1305.New(key[:])
	poly := poly1305.New(&keyx)
	poly.Write(a)
	fmt.Printf("poly hex %x %d\n", poly.Sum(nil),  len(poly.Sum(nil)))
}
