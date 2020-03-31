// This is a jenkins hash 64 
package jenkins64

import (
	"hash"
	"encoding/binary"
	"github.com/tildeleb/hashland/jenkins"
)

// The size of a jenkins checksum in bytes.
const Size = 8 

// The blocksize of jenkins in bytes.
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

	d.hash = jenkins.Hash264(p, d.seed)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the jenkins 64 checksum. 
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

