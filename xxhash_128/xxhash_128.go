// xxhash 128 is a 128-bit or 16 byte hash of the xxhash 
package xxhash_128

import (
	"hash"
	"github.com/OneOfOne/xxhash"
_	"errors"
)

// The size of a xxhash checksum in bytes.
const Size = 16

// The blocksize of xxhash in bytes.
const BlockSize =  32


// digest represents the partial evaluation of a checksum.
type digest struct {
	xxhash *xxhash.XXHash64
	xxhash2 *xxhash.XXHash64
	start    int
	end      int
}

func init() {
}

func (d *digest) Reset() {
	d.xxhash.Reset()
	d.xxhash2.Reset()
}

func (d *digest) Size() int { return Size }
func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {
	d.xxhash.Write(p)
	d.xxhash2.Write(p)
	return

}

// New returns a new xxhash with two hash signature with different keys
func New(params ...uint64) hash.Hash {
	var seed uint64  = 991209123091283
	var seed2 uint64 = 9881223091283

	if (len(params) == 2) {
		seed  = params[0]
		seed2 = params[1]
	}


	d := new(digest)
	// h  := xxhash.NewS64(991209123091283)
	h  := xxhash.NewS64(seed)
	d.xxhash = h

	h2  := xxhash.NewS64(seed2)
	d.xxhash2 = h2

	return d
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.

	return append(d.xxhash.Sum(in), d.xxhash2.Sum(in)...)

}

