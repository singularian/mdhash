// Package t1ha hash
//
// https://github.com/dgryski/go-t1ha
//
// https://github.com/leo-yuriev/t1ha
// 
package t1ha

import (
	"hash"
	"encoding/binary"
	"github.com/dgryski/go-t1ha"
)

// The size of a t1ha checksum in bytes.
const Size = 8

// The blocksize of t1ha in bytes.
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

	d.hash = t1ha.Sum64(p, d.seed)
	binary.BigEndian.PutUint64(d.hashBytes, d.hash)
	return

}

// New returns a new hash.Hash computing the t1ha 64 checksum. 
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

