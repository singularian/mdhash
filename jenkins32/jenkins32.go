// This is a jenkins hash 32 
package jenkins32

import (
	"hash"
	"encoding/binary"
	"github.com/tildeleb/hashland/jenkins"
)

// The size of a jenkins checksum in bytes.
const Size = 4

// The blocksize of jenkins in bytes.
const BlockSize = 4 

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int

	seed uint32
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

	d.hash32 = jenkins.Hash232(p, d.seed)
	binary.BigEndian.PutUint32(d.hashBytes, d.hash32)
	return

}

// New returns a new hash.Hash computing the jenkins 32 checksum. 
func New(start int, end int, seed uint32) hash.Hash {
	d := new(digest)

	d.seed = seed

	r := make([]byte, 4)
	d.hashBytes = r

	d.start = start 
	d.end   = end


	return d
}

// return the sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

