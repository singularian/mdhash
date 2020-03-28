// This is a farm Hash 32 
// this is an example of a write your own signature to add to mdencode
package farmHash32

import (
	"hash"
	"github.com/leemcloughlin/gofarmhash"
_	"errors"
)

// The size of a SHA-1 checksum in bytes.
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
	for i := uint32(0); i < 4; i++ {
		d.hashBytes[i] = byte((d.hash32 >> (8 * i)) & 0xff)
	}
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

func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

