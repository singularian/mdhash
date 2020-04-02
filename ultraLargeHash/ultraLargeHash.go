// This is the ultraLargeHash 
// a large hash generator
// 512 byte experimental non crypo signature generator

package ultraLargeHash 

import (
_	"fmt"
	"hash"
	"math/big"
_	"errors"
)

// The size of a ultraLargeHash checksum in bytes.
const Size = 512

// The blocksize of ultraLargeHash in bytes.
const BlockSize = 512

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int
	hash  *big.Int 
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
	d.hashBytes = d.hashBytes[:]
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

// Write the bytes array and compute the ultraLargeHash hash
func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = big.NewInt(1)
	d.hash = d.ulh(p)
	return

}

// New returns a new hash.Hash computing the ultraLargeHash checksum. 
func New(start int, end int) hash.Hash {
	d := new(digest)

	r := make([]byte, 512)
	d.hashBytes = r

	d.start = start 
	d.end   = end


	return d
}

// return the ultraLargeHash sum
func (d *digest) Sum(in []byte) []byte {
	copy(d.hashBytes[:], d.hash.Bytes())
	return d.hashBytes[d.start:d.end]
}

// compute the djb2 of a byte array
func (d *digest) ulh(s []byte) *big.Int {
        // var hash uint64 = 5381
	// a := big.NewInt(1)
	// b := big.NewInt(1)
	// var hash big.Int
	hash := big.NewInt(1)

	hash = hash.Exp(big.NewInt(2), big.NewInt(4094), nil)

	// fmt.Println("hash bytes ", hash.Bytes(), hash)

        for _, c := range s {
		if int64(c) != 0 {
		hash = hash.Mul(hash, big.NewInt(int64(c)))
		// hash.Mul(hash, big.NewInt(78972222))
		}
        }

	// fmt.Println("hash bytes ", hash.Bytes())
        return hash
}

