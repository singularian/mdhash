// xxhash 128 is a 128-bit or 16 byte hash of the xxhash 
package xxhash_128

import (
	"hash"
	"github.com/OneOfOne/xxhash"
_	"errors"
)

// The size of a xxhash checksum in bytes.
const Size = 16

// The blocksize of SHA-1 in bytes.
const BlockSize = 32 


// digest represents the partial evaluation of a checksum.
type digest struct {
	// xxhash hash.Hash
	// xxhash2 hash.Hash
	xxhash *xxhash.XXHash64
	xxhash2 *xxhash.XXHash64
	// xxhash *xxhash.XXHash64
	// xxhash2 *xxhash.XXHash64
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

// New returns a new hash.Hash computing the SHA1 checksum. The Hash also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
func New(params ...int) hash.Hash {
	d := new(digest)
	h  := xxhash.NewS64(991209123091283)
	d.xxhash = h

	h2  := xxhash.NewS64(9881223091283)
	d.xxhash2 = h2

/*	d.start = 0
	d.end   = 16

	if (len(params) == 2) {
		d.start = params[0]
		d.end   = params[1]
	}
*/
	return d
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d so that caller can keep writing and summing.

	//////////s1 := append(d.xxhash.Sum(in))
	// s2 := append(d.xxhash2.Sum(in))
	// s1 = append(s1, s2...)
	s1 := append(d.xxhash.Sum(in), d.xxhash2.Sum(in)...)

	// fmt.Println("test ", s1, s2)
	// fmt.Println("test ", d.xxhash.Sum(in), d.xxhash2.Sum(in))

	return s1

	// return sl[d.start:d.end]
}

