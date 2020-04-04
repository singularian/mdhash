// This is a marvin 32 
// https://github.com/dgryski/go-marvin32
package marvin32

import (
	"hash"
	"encoding/binary"
	"github.com/dgryski/go-marvin32"
_	"errors"
)

// The size of a marvin 32 checksum in bytes.
const Size = 4

// The blocksize of marvin 32 in bytes.
const BlockSize = 4 

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int

	seed uint64
	hash32 uint32
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash32 = marvin32.Sum32(d.seed, p)
	binary.BigEndian.PutUint32(d.hashBytes, d.hash32)
	return

}

// New returns a new hash.Hash computing the marvin 32 checksum. 
func New(start int, end int, seed uint64) hash.Hash {
	d := new(digest)

	r := make([]byte, 4)
	d.hashBytes = r

	d.seed  = seed

	d.start = start 
	d.end   = end


	return d
}

// return the sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

