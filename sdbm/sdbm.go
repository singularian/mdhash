// This is the sdbm hash 
// 

package sdbm

import (
	"hash"
	"encoding/binary"
_	"errors"
)

// The size of a sdbm checksum in bytes.
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

// Write the bytes array and comput the sdbm hash
func (d *digest) Write(p []byte) (nn int, err error) {

	d.hash = d.sdbm(p)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the sdbm checksum. 
func New(start int, end int) hash.Hash {
	d := new(digest)

	r := make([]byte, 8)
	d.hashBytes = r

	d.start = start 
	d.end   = end


	return d
}

// return the sdbm sum
func (d *digest) Sum(in []byte) []byte {

	return d.hashBytes[d.start:d.end]
}

// compute the sdbm of a byte array
func (d *digest) sdbm(s []byte) uint64 {
        var hash uint64 = 0 

        for _, c := range s {
                hash = uint64(c) + (hash << 6) + (hash << 16) - hash
        }

        return hash
}

