// This is a farm Hash 32 
package farmHash32

import (
	"hash"
	"encoding/binary"
	"github.com/leemcloughlin/gofarmhash"
_	"errors"
)

// The size of a farm hash 32 checksum in bytes.
const Size = 4

// The blocksize of farm hash in bytes.
const BlockSize = 4 

// digest represents the partial evaluation of a checksum.
type digest struct {
	key []byte
	start    int
	end      int
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

	d.hash32 = farmhash.Hash32(p)
	binary.BigEndian.PutUint32(d.hashBytes, d.hash32)
	return

}

// New returns a new hash.Hash computing the farm hash 32 checksum. 
func New(start int, end int, key []byte) hash.Hash {
	d := new(digest)

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

