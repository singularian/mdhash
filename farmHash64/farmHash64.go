// This is a farm Hash 64 
package farmHash64

import (
	"hash"
	"encoding/binary"
	"github.com/leemcloughlin/gofarmhash"
_	"errors"
)

// The size of a farmhash 64 checksum in bytes.
const Size = 8 

// The blocksize of farm hash in bytes.
const BlockSize = 8 

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int
	seed  uint64
	hash  uint64
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

// compute the hash
func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = farmhash.Hash64WithSeed(p, d.seed)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the farm hash 32 checksum. 
func New(start int, end int, seed uint64) hash.Hash {
	d := new(digest)

	r := make([]byte, 8)
	d.hashBytes = r
	d.seed = seed

	d.start = start 
	d.end   = end


	return d
}

// return the sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

