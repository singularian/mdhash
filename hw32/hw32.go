// This is a Highway Hash 32 
// this is an example of a write your own signature to add to mdencode
package hw32

import (
	"hash"
	"github.com/minio/highwayhash"
_	"errors"
)

// The size of a SHA-1 checksum in bytes.
const Size = 4

// The blocksize of SHA-1 in bytes.
const BlockSize = 8

// digest represents the partial evaluation of a checksum.
type digest struct {
	hw64 hash.Hash
	key []byte
	start    int
	end      int
}

func init() {
}

func (d *digest) Reset() {
	d.hw64.Reset()
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {

	_, err = d.hw64.Write(p)
	if (err != nil) {
		panic(err)
	}
	return

}

// New returns a new hash.Hash computing the SHA1 checksum. The Hash also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
func New(start int, end int, key []byte) hash.Hash {
	d := new(digest)
	h, _  := highwayhash.New64(key)
	d.hw64 = h

	d.start = start 
	d.end   = end


	return d
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.
	sl := d.hw64.Sum(in)

	return sl[d.start:d.end]
}

