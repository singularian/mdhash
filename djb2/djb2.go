// This is the djb2 hash 
// this is an example of a write your own signature to add to mdencode
package djb2

import (
	"hash"
	"encoding/binary"
_	"errors"
)

// The size of a djb2 checksum in bytes.
const Size = 8

// The blocksize of farm hash in bytes.
const BlockSize = 8

// digest represents the partial evaluation of a checksum.
type digest struct {
	start    int
	end      int
	hash  uint64
	hashBytes []byte
}

func init() {
}

func (d *digest) Reset() {
	d.hashBytes = d.hashBytes[:]
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = d.djb2(p)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the farm hash 32 checksum. 
func New(start int, end int) hash.Hash {
	d := new(digest)

	r := make([]byte, 8)
	d.hashBytes = r

	d.start = start 
	d.end   = end


	return d
}

// return the dbj2 sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

// compute the djb2 of a byte array
func (d *digest) djb2(s []byte) uint64 {
        var hash uint64 = 5381

        for _, c := range s {
                hash = ((hash << 5) + hash) + uint64(c)
                // the above line is an optimized version of the following line:
                // hash = hash * 33 + int64(c)
                // which is easier to read and understand...
        }

        return hash
}

