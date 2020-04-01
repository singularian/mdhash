// This is a spooky hash 64 
package spooky64

import (
	"hash"
	"encoding/binary"
	"github.com/tildeleb/hashland/spooky"
)

// The size of a spooky checksum in bytes.
const Size = 8 

// The blocksize of spooky in bytes.
const BlockSize = 8 

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int

	seed uint64
	hash uint64
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

// hash a byte slice
func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = spooky.Hash64(p, d.seed)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the spooky 64 checksum. 
func New(start int, end int, seed uint64) hash.Hash {
	d := new(digest)

	d.seed = seed

	r := make([]byte, 8)
	d.hashBytes = r

	d.start = start 
	d.end   = end


	return d
}

// return the sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

